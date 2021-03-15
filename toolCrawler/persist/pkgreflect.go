// Code generated by github.com/newclarity/PackageReflect DO NOT EDIT.

package persist

import "reflect"

var Types = map[string]reflect.Type{
	"Existence": reflect.TypeOf((*Existence)(nil)).Elem(),
	"Hash": reflect.TypeOf((*Hash)(nil)).Elem(),
	"Host": reflect.TypeOf((*Host)(nil)).Elem(),
	"Item": reflect.TypeOf((*Item)(nil)).Elem(),
	"Request": reflect.TypeOf((*Request)(nil)).Elem(),
	"Resource": reflect.TypeOf((*Resource)(nil)).Elem(),
	"SqlId": reflect.TypeOf((*SqlId)(nil)).Elem(),
	"SqlResult": reflect.TypeOf((*SqlResult)(nil)).Elem(),
	"Storage": reflect.TypeOf((*Storage)(nil)).Elem(),
	"Storager": reflect.TypeOf((*Storager)(nil)).Elem(),
	"Visited": reflect.TypeOf((*Visited)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"ErroredPage": reflect.ValueOf(ErroredPage),
	"GetDbFilepath": reflect.ValueOf(GetDbFilepath),
	"GetErroredSubdir": reflect.ValueOf(GetErroredSubdir),
	"GetIndexedSubdir": reflect.ValueOf(GetIndexedSubdir),
	"GetQueuedSubdir": reflect.ValueOf(GetQueuedSubdir),
	"GetQueuedUrls": reflect.ValueOf(GetQueuedUrls),
	"GetSubdir": reflect.ValueOf(GetSubdir),
	"GetSubdirFilepath": reflect.ValueOf(GetSubdirFilepath),
	"GetUrlFilename": reflect.ValueOf(GetUrlFilename),
	"HasQueuedUrls": reflect.ValueOf(HasQueuedUrls),
	"IndexedPage": reflect.ValueOf(IndexedPage),
	"IsSqlUniqueError": reflect.ValueOf(IsSqlUniqueError),
	"NewHash": reflect.ValueOf(NewHash),
	"NewHost": reflect.ValueOf(NewHost),
	"NewResource": reflect.ValueOf(NewResource),
	"Persist": reflect.ValueOf(Persist),
	"WriteFile": reflect.ValueOf(WriteFile),
}

var Variables = map[string]reflect.Value{
	"ErrNonIndexableUrl": reflect.ValueOf(&ErrNonIndexableUrl),
}

var Consts = map[string]reflect.Value{
	"CanExist": reflect.ValueOf(CanExist),
	"CannotExist": reflect.ValueOf(CannotExist),
	"DeleteQueueItemDml": reflect.ValueOf(DeleteQueueItemDml),
	"DeleteQueueItemsByHashDml": reflect.ValueOf(DeleteQueueItemsByHashDml),
	"ErroredDir": reflect.ValueOf(ErroredDir),
	"IndexedDir": reflect.ValueOf(IndexedDir),
	"InsertHostDml": reflect.ValueOf(InsertHostDml),
	"InsertQueueItemDml": reflect.ValueOf(InsertQueueItemDml),
	"InsertResourceDml": reflect.ValueOf(InsertResourceDml),
	"InsertVisitedDml": reflect.ValueOf(InsertVisitedDml),
	"JsonFileTemplate": reflect.ValueOf(JsonFileTemplate),
	"MustExist": reflect.ValueOf(MustExist),
	"QueuedDir": reflect.ValueOf(QueuedDir),
	"RecordAddFailed": reflect.ValueOf(RecordAddFailed),
	"RecordAdded": reflect.ValueOf(RecordAdded),
	"RecordExisted": reflect.ValueOf(RecordExisted),
	"RecordUnknown": reflect.ValueOf(RecordUnknown),
	"SelectHostByIdDml": reflect.ValueOf(SelectHostByIdDml),
	"SelectHostBySDPDml": reflect.ValueOf(SelectHostBySDPDml),
	"SelectHostByUrlDml": reflect.ValueOf(SelectHostByUrlDml),
	"SelectQueueCountDml": reflect.ValueOf(SelectQueueCountDml),
	"SelectQueueItemByHashDml": reflect.ValueOf(SelectQueueItemByHashDml),
	"SelectQueueItemDml": reflect.ValueOf(SelectQueueItemDml),
	"SelectResourceByHashDml": reflect.ValueOf(SelectResourceByHashDml),
	"SelectResourceByIdDml": reflect.ValueOf(SelectResourceByIdDml),
	"SelectResourceCountByIdDml": reflect.ValueOf(SelectResourceCountByIdDml),
	"SelectResourceCountDml": reflect.ValueOf(SelectResourceCountDml),
	"SelectResourceDml": reflect.ValueOf(SelectResourceDml),
	"SelectVisitedStatsByHashDml": reflect.ValueOf(SelectVisitedStatsByHashDml),
	"SqliteDbFilename": reflect.ValueOf(SqliteDbFilename),
}

