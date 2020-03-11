// expect：be sure to finish!
// author：KercyLAN
// create at：2020-2-29 22:51
// todo：待优化，断点续传待实现
package upload

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/KercyLAN/dev-kits/utils/kfile"
	"github.com/KercyLAN/dev-kits/utils/kstreams"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"
)

// 上传失败的文件集合结构描述。
type FailedFiles map[string]error

// 文件上传接收器
//
// 用于对HTTP POST请求中的文件上传操作进行处理，
// 由于文件上传接收器是支持云端应用的，所以所有文件均以UUID的形式在磁盘存储。
//
// 在默认不配置storageDir的情况下，默认的存储路径为"./storage"目录下，而一些临时文件则会存储在"./storage/temp"目录下。
//
// 由于FileUploadReceiver在设计之初便考虑其需要极强的适应性，因此在非特殊场景需求的情况下，通常维护FileUploadReceiver的单例即可。
//
// 在使用FileUploadReceiver过程中，客户端不需要做过多的处理，便可以实现诸如下方的各项功能：
//	- 普通的单个或多个文件上传；
//	- 单个或多个文件的快速上传（秒传）；
//	- 普通上传和快速上传混合使用（如文件1进行普通上传，文件2进行快速上传）；
//	- 普通的文件分块上传；
//	- 快速上传模式下的分块上传；
//	- 文件传输过程劫持篡改检查；
//	- 多文件上传是否要求必须全部上传成功的控制；
//	- 对上传的文件大小进行限制；
type FileUploadReceiver struct {
	*partProcessor                                       // 分块上传处理器
	storageDir               string                      // 文件存储目录
	readMaxMemory            int64                       // ReadForm最大的内存数，会预留10MB
	saveFileBufferSize       int64                       // 每次读取文件缓冲区大小
	allPass                  bool                        // 是否需要所有文件均上传成功才视为上传成功
	fastUploadHandler        func(hash string) *FileMeta // 查询文件是否已上传的处理函数
	completenessCheckHandler func(data []byte) string    // 完整性检查处理函数
	partUploadBufferSize     int64                       // 分块上传每个块的缓冲区大小
	fileMaxSize              int64                       // 上传文件最大大小
	partFileMaxSize          int64                       // 分块上传文件最大大小
}

// 设置分块上传的文件最大允许的内存空间占用大小
//
// 上传文件中，如果这个fileMaxSize大于0，将会对文件大小进行校验。
//
// 当fileMaxSize小于0的时候，将会设置为0。
func (slf *FileUploadReceiver) SetPartFileMaxSize(fileMaxSize int64) {
	if fileMaxSize < 0 {
		fileMaxSize = 0
	}
	slf.partFileMaxSize = fileMaxSize
}

// 设置上传的文件最大允许的内存空间占用大小
//
// 上传文件中，如果这个fileMaxSize大于0，将会对文件大小进行校验。
//
// 当fileMaxSize小于0的时候，将会设置为0。
func (slf *FileUploadReceiver) SetFileMaxSize(fileMaxSize int64) {
	if fileMaxSize < 0 {
		fileMaxSize = 0
	}
	slf.fileMaxSize = fileMaxSize
}

// 设置分块上传每个块的缓冲区大小
//
// 当上传的块占用空间大小小于这个值的情况下，则会直接将其数据存储在内存中以加快效率。
//
// 当上传的块占用空间大小大于这个值的情况下，会将其数据存储到storageDir的临时目录中。
//
// 合理的设置这个值可以防止由于内存申请过大造成的内存溢出问题，
// 同时也可以在性能和稳定性之间保持一个平衡。
func (slf *FileUploadReceiver) SetPartUploadBufferSize(partUploadBufferSize int64) {
	slf.partUploadBufferSize = partUploadBufferSize
}

// 设置文件完整性检查处理函数
//
// 文件完整性检查处理函数用于检查文件是否在上传过程中被篡改。
//
// 使用某种加密方式加密数据后返回一个计算得出的hash值即可用作完整性检查，通常建议CRC、MD5或者SHA256加密。
//
// 使用文件完整性检查的情况下，所消耗的内存会因为要将文件内容存储下来而变为文件占用空间大小 + readMaxMemory的值，效率也会有所降低。
func (slf *FileUploadReceiver) SetCompletenessCheckHandler(handler func(data []byte) string) {
	slf.completenessCheckHandler = handler
}

// 设置快速上传查询函数
//
// 快速上传查询被用于快速上传功能中，所设置的handler需要实现查询传入的hash是否已经在之前被上传过。
// 这个hash通常是由客户端使用和服务端相同的计算方式得到的，服务端在上传文件的时候，如果这个hash没有被上传过，
// 那么应该执行普通的上传流程，并将上传后的结果使用数据库等方式进行存储，而handler便是在存储结果中查阅是否拥
// 有这个记录，并返回一个FileMeta。
//
// 当返回的FileMeta不为空的时候，视为已上传，则进行快速上传操作。
func (slf *FileUploadReceiver) SetFastUploadHandler(handler func(hash string) *FileMeta) {
	slf.fastUploadHandler = handler
}

// 设置是否不允许存在文件上传失败的情况
//
// 在多文件上传的过程中，如果这个值为true，那么当存在一个文件上传失败的情况，则所有文件上传均视为失败。
// 当这个值为false的时候，如果有文件上传失败则会对其进行记录，并在上传结束后将成功上传的文件FileMeta和
// 上传失败的文件信息进行返回。
func (slf *FileUploadReceiver) SetFailureAllowed(allowOr bool) {
	slf.allPass = allowOr
}

// 设置ReadForm最大的内存数量
//
// 默认golang会预留10MB的内存供给非文件的数据。如果不需要这个预留的内存，则在maxMemory中减去10 * 1024 * 1024即可，不过这样做并不建议。
//
// 当所读取的内容大小超出了这个值，则会缓存在临时文件中，效率将会有所降低。
func (slf *FileUploadReceiver) SetReadMaxMemory(maxMemory int64) {
	slf.readMaxMemory = maxMemory
}

// 设置每次读取文件缓冲区大小
//
// 当文件大小大于这个值的时候，会分为多次进行读取。
//
// 当文件大小小于这个值的时候，会一次性进行读取。
//
// 当这个值过大的情况下会导致一次性申请的内存过多，造成内存浪费。
func (slf *FileUploadReceiver) SetSaveFileBufferSize(size int64) {
	slf.saveFileBufferSize = size
}

// 合并分块上传的文件
//
// 在合并的时候partProcessor会对这个hash所标记的合并任务进行检查。
//
// 当所有分块没有完全上传完毕的时候，则会返回一个error，里面会补充一些提示和进度信息。
func (slf *FileUploadReceiver) DisposePartMerge(hash string, fileName string) (*FileMeta, error) {
	var fileMeta *FileMeta
	if !slf.IsFinish(hash) {
		if slf.partProcessor.taskPartCounter[hash] == 0 {
			return fileMeta, errors.New("this file part group task not exist")
		}
		return fileMeta, errors.New(fmt.Sprintf("there are still some blocks that have not been uploaded. quantity completed: %v/%v", len(slf.task[hash]), slf.taskPartCounter[hash]))
	}

	if slf.fastUploadHandler != nil {
		fileMeta := slf.taskFast[hash]
		if fileMeta != nil {
			return fileMeta, nil
		}
	}

	fileMeta = NewFileMeta(fileName, slf.getPartSize(hash))
	file, err := os.Create(fmt.Sprintf("%v\\%v%v", slf.storageDir, fileMeta.id, fileMeta.ext))
	if err != nil {
		return nil, err
	}
	defer file.Close()


	writeData := make([][]byte, 0)
	var eachErr error
	kstreams.EachMapSort(slf.task[hash], func(partIndex int, data []byte) {
		if len(data) == 0 {
			eachErr = kfile.ReadBlockHook(slf.taskTemp[hash][partIndex], int(slf.saveFileBufferSize), func(data []byte) {
				writeData = append(writeData, data)
			})
		}else {
			writeData = append(writeData, data)
		}
	})
	if eachErr != nil {
		return nil, eachErr
	}

	for _, data := range writeData {
		_, err := file.Write(data)
		if err != nil {
			return nil, err
		}
	}

	defer slf.clear(hash)
	fileMeta.hash = hash
	return fileMeta, nil
}

// 处理分块文件上传请求
//
// 分块上传时FileUploadReceiver会读取请求中的一个文件和一个文本内容来作为块进行上传。
//
// 在HTML中，无需为file input以及hidden input设置特定的name属性，FileUploadReceiver对此不做任何约束。
//
// 在使用分块上传的时候应该注意一下几点：
//	- 分块上传依赖于pratProcessor，在即将开始一个分块上传任务之前，应当调用FileUploadReceiver.NewPartTask来创建一个分块任务；
//	- 分块上传应当为多次请求来进行每个块的上传，比如三个块则发起三次请求；
//	- 在分块上传中，每个请求仅允许存在一个文件和一个文本，而且文本的格式必须符合“hash(string);sort(int):size(int64)”，这个文本表示了最终得到的文件hash应该是什么，这个块在这个文件中的什么位置，这个最终文件的文件大小应该是多大；
//	- 分块上传的每一块文件上传均支持快速上传的特性，客户端无需做特殊处理，服务端仅需要使用过FileUploadReceiver.SetFastUploadHandler即可；
//
// 不用担心会因为多次请求而造成每个块分散开来，FileUploadReceiver.NewPartTask会为他们之间建立关系；
//
// 为了避免内存泄漏，在每一块文件大小超过一个特定的阈值的时候，则这个分块文件会被保存在本地。
//
// 当一个文件的分块即存在大于这个阈值和小于这个阈值的情况，他们会分散在内存和磁盘中，最后会进行合并。
//
// 如果需要调整这个值，可以使用FileUploadReceiver.SetPartUploadBufferSize函数来进行配置。
func (slf *FileUploadReceiver) DisposePart(request *http.Request) (string, error) {
	multipartReader, err := initRequestPart(request); if err != nil {
		return "", err
	}

	fileHeader, hash, fileSize, partIndex, err := readFormPart(multipartReader, slf.readMaxMemory)
	if err != nil {
		return hash, err
	}
	if slf.partFileMaxSize > fileSize {
		return hash, errors.New("file size out range of")
	}
	if slf.partProcessor.task[hash] == nil {
		return hash, errors.New("file part upload must use function \"FileUploadReceiver.NewPartTask\" create a file group task")
	}

	if slf.fastUploadHandler != nil {
		slf.taskFast[hash] = slf.fastUploadHandler(hash)
		if slf.taskFast[hash] != nil {
			slf.taskFast[hash].uploadAt = time.Now()
			slf.taskFast[hash].isFastUpload = true
			if err := slf.addPart(hash, partIndex, make([]byte, 0), slf); err != nil {
				return hash, err
			}
			return hash, nil
		}
	}

	formFile, err := fileHeader.Open()
	defer formFile.Close()
	if err != nil {
		return hash, err
	}

	if fileHeader.Size > slf.partUploadBufferSize {
		filePath := fmt.Sprintf("%v/%v/%v.part.temp", slf.storageDir, "temp", uuid.NewV4().String())
		file, err := os.Create(filePath)
		if err != nil {
			return hash, err
		}
		defer file.Close()

		nowLen := int64(0)
		for {
			buffer := make([]byte, slf.saveFileBufferSize)
			readLen, readErr := formFile.ReadAt(buffer, nowLen)
			nowLen += int64(readLen)

			_, err = file.Write(buffer[0:readLen])
			if err != nil {
				return hash, err
			}
			if readErr != nil {
				if readErr == io.EOF {
					break
				}
				return hash, readErr
			}
		}
		if err := slf.addPartPath(hash, partIndex, filePath); err != nil {
			return hash, err
		}
	}else {
		fileBytes := make([][]byte, 0)
		nowLen := int64(0)
		for {
			buffer := make([]byte, slf.saveFileBufferSize)
			readLen, err := formFile.ReadAt(buffer, nowLen)
			nowLen += int64(readLen)
			fileBytes = append(fileBytes, buffer[0:nowLen])
			if err != nil {
				if err == io.EOF {
					break
				}
				return hash, err
			}
		}
		if err := slf.addPart(hash, partIndex, bytes.Join(fileBytes, make([]byte, 0)), slf); err != nil {
			return hash, err
		}
	}

	return hash, nil
}

// 处理文件上传请求
//
// 在使用上传处理功能的时候应该注意以下两点：
//	- fileGroup表示如“<input name="group_name" type="file">”这个html元素中的name属性，相同name属性的视为同一个fileGroup；
//	- valueGroup表示如“<input name="group_name" type="hidden" value="1">”这个html元素中的name属性，相同name属性的视为同一个valueGroup。
//
// FileUploadReceiver在处理请求时候会有几种情况：
//	- 情况1：当获取到的valueGroup不包含任何成员的时候，表明为普通的文件上传；
//	- 情况2：当获取到的fileGroup和valueGroup包含1个成员的时候，表明为快速上传，即秒传；
//		·同时，在满足情况2和已经设置过SetCompletenessCheckHandler的情况下，会对文件的完整性进行检查。
//	- 情况3：当获取到的fileGroup和valueGroup的成员数不是1:1的时候，会返回error。
//
// 秒传情况下的请求中，保证每个文件的name属性和表示其hash的POST参数name属性相同。
//
// 应该保证如下方这样的form结构。
//	<form enctype="multipart/form-data" action="_URL_" method=POST>
//		<input name="file1" type="file">
// 		<input name="file1" type="hidden" value="file1-hash">
// 		<input name="file2" type="file">
// 		<input name="file2" type="hidden" value="file2-hash">
//	</form>
func (slf *FileUploadReceiver) Dispose(request *http.Request) ([]*FileMeta, FailedFiles, error) {
	failFiles := make(map[string]error)
	fileMetas, multipartReader, err := initRequest(request); if err != nil {
		return fileMetas, failFiles, err
	}

	fileGroup, valueGroup, err := readForm(multipartReader, slf.readMaxMemory)
	if err != nil {
		return fileMetas, failFiles, err
	}

	err = fileGroup.EachAll(func(groupName string, groupIndex int, fileHeader *multipart.FileHeader) error {
		if valueGroup.Exist(groupName) {
			if valueGroup.GroupLen(groupName) != 1 || fileGroup.GroupLen(groupName) != 1 {
				return errors.New(errMultipleValue)
			}
			// 是否超出限制的大小
			if slf.fileMaxSize > fileHeader.Size {
				return errors.New("file size out of range")
			}
			// 快速上传模式
			if slf.fastUploadHandler != nil {
				fileMeta := slf.fastUploadHandler(valueGroup[groupName][0])
				if fileMeta != nil {
					fileMeta.name = fileHeader.Filename
					fileMeta.ext = path.Ext(fileMeta.name)
					fileMeta.uploadAt = time.Now()
					fileMeta.isFastUpload = true
					fileMetas = append(fileMetas, fileMeta)
					return nil
				}
			}else {
				log.Println("----------------------------------------------------------------")
				log.Println("Warn!!! FileUploadReceiver not use func \"SetFastUploadHandler()\". This will cause the quick upload feature to fail")
				log.Println("----------------------------------------------------------------")
			}
		}

		fileMeta, fileHash, err := saveFile(slf, fileHeader, slf.storageDir, slf.saveFileBufferSize)
		if err != nil {
			return err
		}

		if valueGroup.Exist(groupName) {
			if valueGroup.GroupLen(groupName) == 1 {
				if fileHash == valueGroup[groupName][0] {
					fileMeta.hash = valueGroup[groupName][0]
				}else {
					return errors.New("file hash value exception. the document has been hijacked and tampered with")
				}
			}
		}
		fileMetas = append(fileMetas, fileMeta)
		return nil

	}, func(groupName string, groupIndex int, fileName string, fileSize int64, err error) bool {
		if slf.allPass {
			failFiles = make(map[string]error)
			for gName, headers := range fileGroup {
				for _, header := range headers {
					if gName == groupName && fileName == header.Filename {
						failFiles[fileName] = errors.New(fmt.Sprintf("form name \"%v\" file \"%v\" upload failed. %v", groupName, fileName, err.Error()))
					}else {
						failFiles[fileName] = errors.New(fmt.Sprintf("form name \"%v\" file \"%v\" upload failed. exist failed upload", gName, header.Filename))
					}
				}
			}
			return true
		}
		failFiles[fileName] = errors.New(fmt.Sprintf("form name \"%v\" file \"%v\" upload failed. %v", groupName, fileName, err.Error()))
		return false
	})

	// 当err内容不为errMultipleValue，表示开启了必须全部上传成功的设置
	if err != nil {
		if err.Error() != errMultipleValue {
			return make([]*FileMeta, 0), failFiles, err
		}
		if err.Error() == errMultipleValue {
			return fileMetas, failFiles, err
		}
	}
	return fileMetas, failFiles, nil
}

// 构建一个FileUploadReceiver实例
//
// 当storageDir未设置的情况下，FileUploadReceiver会将所有的文件存储在默认的“./storage”目录下。
//
// 如果需要改变这一情况，可以在构建FileUploadReceiver的时候传入一个目录路径作为参数。
// 如果这个目录不存在的话，FileUploadReceiver会自行进行创建，如果参数存在多个的话，那么只取第一个。
//
// 注意：存储文件需要依赖于节点信息，应该在main包的init函数下调用“application.Initialize()”函数进行应用程序初始化。
func New(storageDir ...string) (*FileUploadReceiver, error) {
	storage := "./stroage"
	if len(storageDir) > 0 {
		storage = storageDir[0]
	}

	this := &FileUploadReceiver{
		storageDir: storage,
		readMaxMemory: 1024,
		saveFileBufferSize: 5 * 1024 * 1024,
		partUploadBufferSize: 30 * 1024 * 1024,
		partProcessor: &partProcessor{
			taskFast: map[string]*FileMeta{},
			task: map[string]map[int][]byte{},
			taskPartCounter: map[string]int{},
			taskTemp: map[string]map[int]string{},
		},
	}
	if err := os.MkdirAll(this.storageDir, os.ModeDir); err != nil {
		return nil, errors.New("the storageDir directory could not be initialized\r\n" + err.Error())
	}
	if err := os.MkdirAll(this.storageDir + "/temp", os.ModeDir); err != nil {
		return nil, errors.New("the storageDir directory could not be initialized\r\n" + err.Error())
	}
	return this, nil
}
