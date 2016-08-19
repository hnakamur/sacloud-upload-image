package api

/************************************************
  generated by IDE. for [DiskAPI]
************************************************/

import (
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

/************************************************
   To support influent interface for Find()
************************************************/

func (api *DiskAPI) Reset() *DiskAPI {
	api.reset()
	return api
}

func (api *DiskAPI) Offset(offset int) *DiskAPI {
	api.offset(offset)
	return api
}

func (api *DiskAPI) Limit(limit int) *DiskAPI {
	api.limit(limit)
	return api
}

func (api *DiskAPI) Include(key string) *DiskAPI {
	api.include(key)
	return api
}

func (api *DiskAPI) Exclude(key string) *DiskAPI {
	api.exclude(key)
	return api
}

func (api *DiskAPI) FilterBy(key string, value interface{}) *DiskAPI {
	api.filterBy(key, value, false)
	return api
}

// func (api *DiskAPI) FilterMultiBy(key string, value interface{}) *DiskAPI {
// 	api.filterBy(key, value, true)
// 	return api
// }

func (api *DiskAPI) WithNameLike(name string) *DiskAPI {
	return api.FilterBy("Name", name)
}

func (api *DiskAPI) WithTag(tag string) *DiskAPI {
	return api.FilterBy("Tags.Name", tag)
}
func (api *DiskAPI) WithTags(tags []string) *DiskAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

func (api *DiskAPI) WithSizeGib(size int) *DiskAPI {
	api.FilterBy("SizeMB", size*1024)
	return api
}

func (api *DiskAPI) SortBy(key string, reverse bool) *DiskAPI {
	api.sortBy(key, reverse)
	return api
}

func (api *DiskAPI) SortByName(reverse bool) *DiskAPI {
	api.sortByName(reverse)
	return api
}

func (api *DiskAPI) SortBySize(reverse bool) *DiskAPI {
	api.sortBy("SizeMB", reverse)
	return api
}

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

func (api *DiskAPI) New() *sacloud.Disk {
	return sacloud.CreateNewDisk()
}

//func (api *DiskAPI) Create(value *sacloud.Disk) (*sacloud.Disk, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.create(api.createRequest(value), res)
//	})
//}

func (api *DiskAPI) Read(id string) (*sacloud.Disk, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

func (api *DiskAPI) Update(id string, value *sacloud.Disk) (*sacloud.Disk, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *DiskAPI) Delete(id string) (*sacloud.Disk, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

/************************************************
  Inner functions
************************************************/

func (api *DiskAPI) setStateValue(setFunc func(*sacloud.Request)) *DiskAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *DiskAPI) request(f func(*sacloud.Response) error) (*sacloud.Disk, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Disk, nil
}

func (api *DiskAPI) createRequest(value *sacloud.Disk) *sacloud.Request {
	req := &sacloud.Request{}
	req.Disk = value
	return req
}
