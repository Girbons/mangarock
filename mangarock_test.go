package mangarock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var client = NewClient()
var options = map[string]string{"country": "italy"}

func TestMangaRockInfo(t *testing.T) {
	client.SetOptions(options)
	result, _ := client.Info("mrs-serie-35593")

	assert.Equal(t, 0, result.Code)
	assert.Equal(t, "Boruto: Naruto Next Generations", result.Data.Name)
	assert.Equal(t, "Ukyo Kodachi", result.Data.Author)
	assert.Equal(t, "https://f01.mrcdn.info/file/mrportal/i/5/8/3/G3.6PwgFb_B.jpg", result.Data.Thumbnail)
	assert.Equal(t, 2, len(result.Data.Authors))
	// check that Author is filled
	assert.Equal(t, "Ukyo Kodachi", result.Data.Authors[0].Name)
	assert.Equal(t, "mrs-author-306911", result.Data.Authors[0].Oid)
	assert.Equal(t, "story", result.Data.Authors[0].Role)
	assert.Equal(t, "https://f01.mrcdn.info/file/mrportal/i/5/7/g/ej.vP9TUgn.jpg", result.Data.Authors[0].Thumbnail)
	assert.Equal(t, 26547669, result.Data.Chapters[0].Cid)
	// check that chapter is filled
	assert.Equal(t, "Vol.1 Number 1: Uzumaki Boruto!!", result.Data.Chapters[0].Name)
	assert.Equal(t, "mrs-chapter-35594", result.Data.Chapters[0].Oid)
	assert.Equal(t, 0, result.Data.Chapters[0].Order)
	// check that Character is filled
	assert.Equal(t, "Hinata Hyuuga", result.Data.Characters[0].Name)
	assert.Equal(t, "mrs-character-311689", result.Data.Characters[0].Oid)
	assert.Equal(t, "https://f01.mrcdn.info/file/mrportal/h/3/3/6/Gr.1P-4Ci00.jpg", result.Data.Characters[0].Thumbnail)
	assert.Equal(t, 8, len(result.Data.Categories))
	assert.Equal(t, 29, len(result.Data.Chapters))
	assert.Equal(t, false, result.Data.Completed)
	assert.Equal(t, "https://f01.mrcdn.info/file/mrportal/h/c/3/0/J_.h_1FHZfW.jpg", result.Data.Cover)
	assert.Equal(t, 1, result.Data.Direction)
	assert.Equal(t, "mrs-serie-35593", result.Data.Oid)
	assert.Equal(t, 529, result.Data.Rank)
	assert.Equal(t, false, result.Data.Removed)
	assert.Equal(t, "Viz", result.Data.Extra["English Publisher"])
	assert.Equal(t, "Shueisha ", result.Data.Extra["Original Publisher"])
	assert.Equal(t, "May 9, 2016 ", result.Data.Extra["Published"])
	assert.Equal(t, 29, result.Data.TotalChapters)
	assert.Equal(t, 553712, result.Data.Mid)
	assert.Equal(t, 71, result.Data.MsID)
}

func TestMangaRockPages(t *testing.T) {
	client.SetOptions(options)

	result, _ := client.Pages("mrs-chapter-100051049")

	assert.Equal(t, 0, result.Code)
	assert.Equal(t, 49, len(result.Data))
}
