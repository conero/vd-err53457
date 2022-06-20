package repo

import (
	"encoding/xml"
)

type svnInfoEnterXml struct {
	XMLName  xml.Name `xml:"entry"`
	Revision string   `xml:"revision,attr"`
	Uuid     string   `xml:"repository>uuid"`
	Author   string   `xml:"commit>author"`
	Date     string   `xml:"commit>date"`
}

type svnInfoXml struct {
	XMLName xml.Name `xml:"info"`
	Enter   svnInfoEnterXml
}

type svnPathXml struct {
	XMLName  xml.Name `xml:"path"`
	Item     string   `xml:"item,attr"`
	Kind     string   `xml:"kind,attr"`
	Filename string   `xml:",innerxml"`
}

// `svn diff`
type svnDiffXml struct {
	XMLName xml.Name     `xml:"diff"`
	Paths   []svnPathXml `xml:"paths>path"`
}

type Svn struct {
	vUrl string
}

func (c *Svn) BaseUrl() string {
	return c.vUrl
}

func (c *Svn) getXmlString(vCmd string) string {
	return ""
}

func (c *Svn) Latest() (RevisionInfo, error) {
	ri := RevisionInfo{}
	return ri, nil
}

func (c *Svn) Patch(from, to string) ([]DiffPath, error) {
	return nil, nil
}
