//
// Copyright 2014, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack43

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type AddUcsManagerParams struct {
	p map[string]interface{}
}

func (p *AddUcsManagerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddUcsManagerParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *AddUcsManagerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddUcsManagerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddUcsManagerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

func (p *AddUcsManagerParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddUcsManagerParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewAddUcsManagerParams(password string, url string, username string, zoneid string) *AddUcsManagerParams {
	p := &AddUcsManagerParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["url"] = url
	p.p["username"] = username
	p.p["zoneid"] = zoneid
	return p
}

// Adds a Ucs manager
func (s *UCSService) AddUcsManager(p *AddUcsManagerParams) (*AddUcsManagerResponse, error) {
	resp, err := s.cs.newRequest("addUcsManager", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddUcsManagerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddUcsManagerResponse struct {
	Name   string `json:"name,omitempty"`
	Url    string `json:"url,omitempty"`
	Zoneid string `json:"zoneid,omitempty"`
	Id     string `json:"id,omitempty"`
}

type ListUcsManagersParams struct {
	p map[string]interface{}
}

func (p *ListUcsManagersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListUcsManagersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListUcsManagersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListUcsManagersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListUcsManagersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListUcsManagersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListUcsManagersParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsManagersParams() *ListUcsManagersParams {
	p := &ListUcsManagersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UCSService) GetUcsManagerID(keyword string) (string, error) {
	p := &ListUcsManagersParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListUcsManagers(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.UcsManagers[0].Id, nil
}

// List ucs manager
func (s *UCSService) ListUcsManagers(p *ListUcsManagersParams) (*ListUcsManagersResponse, error) {
	resp, err := s.cs.newRequest("listUcsManagers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsManagersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListUcsManagersResponse struct {
	Count       int           `json:"count"`
	UcsManagers []*UcsManager `json:"ucsmanager"`
}

type UcsManager struct {
	Zoneid string `json:"zoneid,omitempty"`
	Id     string `json:"id,omitempty"`
	Url    string `json:"url,omitempty"`
	Name   string `json:"name,omitempty"`
}

type ListUcsProfilesParams struct {
	p map[string]interface{}
}

func (p *ListUcsProfilesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *ListUcsProfilesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListUcsProfilesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListUcsProfilesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListUcsProfilesParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
	return
}

// You should always use this function to get a new ListUcsProfilesParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsProfilesParams(ucsmanagerid string) *ListUcsProfilesParams {
	p := &ListUcsProfilesParams{}
	p.p = make(map[string]interface{})
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// List profile in ucs manager
func (s *UCSService) ListUcsProfiles(p *ListUcsProfilesParams) (*ListUcsProfilesResponse, error) {
	resp, err := s.cs.newRequest("listUcsProfiles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsProfilesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListUcsProfilesResponse struct {
	Count       int           `json:"count"`
	UcsProfiles []*UcsProfile `json:"ucsprofile"`
}

type UcsProfile struct {
	Ucsdn string `json:"ucsdn,omitempty"`
}

type ListUcsBladesParams struct {
	p map[string]interface{}
}

func (p *ListUcsBladesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *ListUcsBladesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListUcsBladesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListUcsBladesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListUcsBladesParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
	return
}

// You should always use this function to get a new ListUcsBladesParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsBladesParams(ucsmanagerid string) *ListUcsBladesParams {
	p := &ListUcsBladesParams{}
	p.p = make(map[string]interface{})
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UCSService) GetUcsBladeID(keyword string, ucsmanagerid string) (string, error) {
	p := &ListUcsBladesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["ucsmanagerid"] = ucsmanagerid

	l, err := s.ListUcsBlades(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.UcsBlades[0].Id, nil
}

// List ucs blades
func (s *UCSService) ListUcsBlades(p *ListUcsBladesParams) (*ListUcsBladesResponse, error) {
	resp, err := s.cs.newRequest("listUcsBlades", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsBladesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListUcsBladesResponse struct {
	Count     int         `json:"count"`
	UcsBlades []*UcsBlade `json:"ucsblade"`
}

type UcsBlade struct {
	Ucsmanagerid string `json:"ucsmanagerid,omitempty"`
	Id           string `json:"id,omitempty"`
	Hostid       string `json:"hostid,omitempty"`
	Profiledn    string `json:"profiledn,omitempty"`
	Bladedn      string `json:"bladedn,omitempty"`
}

type AssociateUcsProfileToBladeParams struct {
	p map[string]interface{}
}

func (p *AssociateUcsProfileToBladeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bladeid"]; found {
		u.Set("bladeid", v.(string))
	}
	if v, found := p.p["profiledn"]; found {
		u.Set("profiledn", v.(string))
	}
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *AssociateUcsProfileToBladeParams) SetBladeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bladeid"] = v
	return
}

func (p *AssociateUcsProfileToBladeParams) SetProfiledn(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["profiledn"] = v
	return
}

func (p *AssociateUcsProfileToBladeParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
	return
}

// You should always use this function to get a new AssociateUcsProfileToBladeParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewAssociateUcsProfileToBladeParams(bladeid string, profiledn string, ucsmanagerid string) *AssociateUcsProfileToBladeParams {
	p := &AssociateUcsProfileToBladeParams{}
	p.p = make(map[string]interface{})
	p.p["bladeid"] = bladeid
	p.p["profiledn"] = profiledn
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// associate a profile to a blade
func (s *UCSService) AssociateUcsProfileToBlade(p *AssociateUcsProfileToBladeParams) (*AssociateUcsProfileToBladeResponse, error) {
	resp, err := s.cs.newRequest("associateUcsProfileToBlade", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssociateUcsProfileToBladeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r AssociateUcsProfileToBladeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AssociateUcsProfileToBladeResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Profiledn    string `json:"profiledn,omitempty"`
	Bladedn      string `json:"bladedn,omitempty"`
	Hostid       string `json:"hostid,omitempty"`
	Ucsmanagerid string `json:"ucsmanagerid,omitempty"`
	Id           string `json:"id,omitempty"`
}

type DeleteUcsManagerParams struct {
	p map[string]interface{}
}

func (p *DeleteUcsManagerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *DeleteUcsManagerParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
	return
}

// You should always use this function to get a new DeleteUcsManagerParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewDeleteUcsManagerParams(ucsmanagerid string) *DeleteUcsManagerParams {
	p := &DeleteUcsManagerParams{}
	p.p = make(map[string]interface{})
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// Delete a Ucs manager
func (s *UCSService) DeleteUcsManager(p *DeleteUcsManagerParams) (*DeleteUcsManagerResponse, error) {
	resp, err := s.cs.newRequest("deleteUcsManager", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUcsManagerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteUcsManagerResponse struct {
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type DisassociateUcsProfileFromBladeParams struct {
	p map[string]interface{}
}

func (p *DisassociateUcsProfileFromBladeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bladeid"]; found {
		u.Set("bladeid", v.(string))
	}
	if v, found := p.p["deleteprofile"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("deleteprofile", vv)
	}
	return u
}

func (p *DisassociateUcsProfileFromBladeParams) SetBladeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bladeid"] = v
	return
}

func (p *DisassociateUcsProfileFromBladeParams) SetDeleteprofile(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deleteprofile"] = v
	return
}

// You should always use this function to get a new DisassociateUcsProfileFromBladeParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewDisassociateUcsProfileFromBladeParams(bladeid string) *DisassociateUcsProfileFromBladeParams {
	p := &DisassociateUcsProfileFromBladeParams{}
	p.p = make(map[string]interface{})
	p.p["bladeid"] = bladeid
	return p
}

// disassociate a profile from a blade
func (s *UCSService) DisassociateUcsProfileFromBlade(p *DisassociateUcsProfileFromBladeParams) (*DisassociateUcsProfileFromBladeResponse, error) {
	resp, err := s.cs.newRequest("disassociateUcsProfileFromBlade", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisassociateUcsProfileFromBladeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r DisassociateUcsProfileFromBladeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DisassociateUcsProfileFromBladeResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Ucsmanagerid string `json:"ucsmanagerid,omitempty"`
	Bladedn      string `json:"bladedn,omitempty"`
	Hostid       string `json:"hostid,omitempty"`
	Profiledn    string `json:"profiledn,omitempty"`
	Id           string `json:"id,omitempty"`
}

type RefreshUcsBladesParams struct {
	p map[string]interface{}
}

func (p *RefreshUcsBladesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *RefreshUcsBladesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *RefreshUcsBladesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *RefreshUcsBladesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *RefreshUcsBladesParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
	return
}

// You should always use this function to get a new RefreshUcsBladesParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewRefreshUcsBladesParams(ucsmanagerid string) *RefreshUcsBladesParams {
	p := &RefreshUcsBladesParams{}
	p.p = make(map[string]interface{})
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// refresh ucs blades to sync with UCS manager
func (s *UCSService) RefreshUcsBlades(p *RefreshUcsBladesParams) (*RefreshUcsBladesResponse, error) {
	resp, err := s.cs.newRequest("refreshUcsBlades", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RefreshUcsBladesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type RefreshUcsBladesResponse struct {
	Profiledn    string `json:"profiledn,omitempty"`
	Ucsmanagerid string `json:"ucsmanagerid,omitempty"`
	Hostid       string `json:"hostid,omitempty"`
	Bladedn      string `json:"bladedn,omitempty"`
	Id           string `json:"id,omitempty"`
}