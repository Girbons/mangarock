package mangarock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var options = map[string]string{"country": "italy"}
var client = NewClientWithOptions(options)

func TestMangaRockInfo(t *testing.T) {
	result, _ := client.Info("mrs-serie-173467")

	assert.Equal(t, 0, result.Code)
	assert.Equal(t, "Naruto", result.Data.Name)
	assert.Equal(t, "Masashi Kishimoto", result.Data.Author)
	assert.Equal(t, "https://f01.mrcdn.info/file/mrportal/j/3/8/3/nE.ffeWKPZd.jpg", result.Data.Thumbnail)
	assert.Equal(t, 2, len(result.Data.Authors))
	// check that Author is filled
	assert.Equal(t, "Masashi Kishimoto", result.Data.Authors[0].Name)
	assert.Equal(t, "mrs-author-324987", result.Data.Authors[0].OID)
	assert.Equal(t, "art", result.Data.Authors[0].Role)
	assert.Equal(t, "https://f01.mrcdn.info/file/mrportal/i/4/c/b/nR.dAgpfbmr.png", result.Data.Authors[0].Thumbnail)
	assert.Equal(t, 26144172, result.Data.Chapters[0].CID)
	// check that chapter is filled
	assert.Equal(t, "Vol.1 Chapter 1: Uzumaki Naruto!", result.Data.Chapters[0].Name)
	assert.Equal(t, "mrs-chapter-173469", result.Data.Chapters[0].OID)
	assert.Equal(t, 0, result.Data.Chapters[0].Order)
	assert.Equal(t, 8, len(result.Data.Categories))
	assert.Equal(t, 700, len(result.Data.Chapters))
	assert.Equal(t, true, result.Data.Completed)
	assert.Equal(t, "https://vgfiles.nabstudio.com/portal/ffc64df55ae3727de601cfa316490c61_173467_cover.jpg", result.Data.Cover)
	assert.Equal(t, 1, result.Data.Direction)
	assert.Equal(t, "mrs-serie-173467", result.Data.OID)
	assert.Equal(t, false, result.Data.Removed)
	assert.Equal(t, "Viz", result.Data.Extra["English Publisher"])
	assert.Equal(t, "Shueisha ", result.Data.Extra["Original Publisher"])
	assert.Equal(t, "Sep 21, 1999 to Nov 10, 2014", result.Data.Extra["Published"])
	assert.Equal(t, 700, result.Data.TotalChapters)
	assert.Equal(t, 545007, result.Data.Mid)
	assert.Equal(t, 71, result.Data.MsID)
}

func TestMangaRockPages(t *testing.T) {
	client.SetOptions(options)

	result, _ := client.Pages("mrs-chapter-100051049")

	assert.Equal(t, 0, result.Code)
	assert.Equal(t, 49, len(result.Data))
}
