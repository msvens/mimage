package metadata

import (
	"io/ioutil"
	"testing"
	"time"
)

const AssetPath = "../assets/"
const LeicaImg = AssetPath + "leica.jpg"
const NikonImg = AssetPath + "nikon.jpg"
const CanonImg = AssetPath + "canon.jpg"

func getAssetBytes(fname string) ([]byte, error) {
	return ioutil.ReadFile(AssetPath + fname)
}

func getEditor(fname string) (*MetaDataEditor, error) {
	if b, e := getAssetBytes(fname); e != nil {
		return nil, e
	} else {
		return NewMetaDataEditor(b)
	}
}

func TestMetaDataEditor_Bytes(t *testing.T) {
	var b []byte
	mde, err := NewMetaDataEditorFile(LeicaImg)
	if err != nil {
		t.Fatalf("Could not open editor: %s", err.Error())
	}

	//Test 1: Just write bytes and read them back again
	if b, err = mde.Bytes(); err != nil {
		t.Errorf(err.Error())
	} else {
		mde, err = NewMetaDataEditor(b)
		if err != nil {
			t.Errorf("Could not open editor from written bytes: %s", err.Error())
		}

	}
	//Test 2: Write bytes after an edit
	err = mde.SetExifDate(ModifyDate, time.Now())
	if b, err = mde.Bytes(); err != nil {
		t.Errorf(err.Error())
	} else {
		mde, err = NewMetaDataEditor(b)
		if err != nil {
			t.Errorf("Could not open editor from written bytes: %s", err.Error())
		}
	}

}

func TestMetaDataEditor_CopyMetaData(t *testing.T) {

}
