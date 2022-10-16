package metadata

import (
	"reflect"
	"testing"
	"time"
)

func TestNewIptcEditor(t *testing.T) {
	images := []string{LeicaImg, NoExifImg}
	for _, fname := range images {
		b := getAssetBytes(fname, t)
		if sl, err := parseJpegBytes(b); err != nil {
			t.Errorf("Could not parse jpeg file %s", fname)
		} else {
			ie, e1 := NewIptcEditor(sl)
			if e1 != nil {
				t.Errorf("Could not open iptc editor got %v", e1)
			}
			if ie.dirty {
				t.Errorf("expected dirty to be false")
			}
			if ie.raw == nil {
				t.Errorf("expected rawIptc to be non nil")
			}
		}
	}
	if _, err := NewIptcEditor(nil); err == nil {
		t.Errorf("Expected error when creating a NewIptcEditor with a nil segment list")
	}
}

func TestNewIptcEditorEmpty(t *testing.T) {
	ie := NewIptcEditorEmpty(false)
	if ie.IsDirty() {
		t.Errorf("Expected non dirty iptc editor")
	}
	if ie.raw == nil {
		t.Errorf("Expected raw to be non nil")
	}
	ie = NewIptcEditorEmpty(true)
	if !ie.IsDirty() {
		t.Errorf("Expected dirty iptc editor")
	}

}

func TestIptcEditor_Bytes(t *testing.T) {
	tests := []string{LeicaImg, NoExifImg}
	for _, fname := range tests {
		je := getJpegEditor(fname, t)
		je.Iptc().SetDirty() //triggers a call to bytes when retrieving the MetaData
		md := jpegEditorMD(je, t)
		actual := uint16(0)
		if err := md.Iptc().ScanEnvelope(IPTCEnvelope_EnvelopeRecordVersion, &actual); err != nil {
			t.Errorf("Expected EnvelopRecordVersion got error: %v", err)
		} else if actual != uint16(4) {
			t.Errorf("Expected EnvelopRecordVersion 4 got %v", actual)
		}
		if err := md.Iptc().ScanEnvelope(IPTCEnvelope_FileFormat, &actual); err != nil {
			t.Errorf("Expected EnvelopFileFormat got error: %v", err)
		} else if actual != uint16(11) {
			t.Errorf("Expected FileFormat 11 got %v", actual)
		}
	}
}

func TestIptcEditor_Clear(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	je.Iptc().Clear(true)
	md := jpegEditorMD(je, t)
	//Expect 4 IPTC tags
	if len(md.Iptc().RawIptc()) != 4 {
		t.Errorf("Expected  mandatory tags got %v", len(md.Iptc().RawIptc()))
	}

}

func TestIptcEditor_IsDirty(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	if je.Iptc().IsDirty() {
		t.Errorf("Expected not dirty iptc editor")
	}
	if err := je.Iptc().SetTitle("some new title"); err != nil {
		t.Errorf("Could not set title")
	}
	if !je.Iptc().IsDirty() {
		t.Errorf("Expected dirty iptc editor after setting title")
	}
	_, _ = je.Bytes()
	if je.Iptc().IsDirty() {
		t.Errorf("Expected not dirty iptc editor after retrieving bytes")
	}
	je.Iptc().SetDirty()
	if !je.Iptc().IsDirty() {
		t.Errorf("Expected dirty iptc editor after forcing dirty")
	}

}

func TestIptcEditor_IsEmpty(t *testing.T) {
	je := getJpegEditor(NoExifImg, t)
	if !je.Iptc().IsEmpty() {
		t.Errorf("Expected editor to be empty")
	}
	if err := je.Iptc().SetTitle("some title"); err != nil {
		t.Errorf("Could not set title")
	}
	if je.Iptc().IsEmpty() {
		t.Errorf("Expected editor to not be empty")
	}
}

func TestIptcEditor_SetApplication(t *testing.T) {
	expEditorialUpdate := "01"
	for _, fname := range []string{LeicaImg, NoExifImg} {
		je := getJpegEditor(fname, t)
		if err := je.Iptc().setApplication(IPTCApplication_EditorialUpdate, expEditorialUpdate); err != nil {
			t.Fatalf("Could not set IPTCApplication_EditorialUpdate for image %s got err %v", fname, err)
		}
		md := jpegEditorMD(je, t)
		actEditorialUpdate := ""
		if err := md.Iptc().ScanApplication(IPTCApplication_EditorialUpdate, &actEditorialUpdate); err != nil {
			t.Fatalf("Could not scan IPTCApplicatin_EditorialUpdate for image %s got err %v", fname, err)
		}
		if expEditorialUpdate != actEditorialUpdate {
			t.Errorf("Expected %v got %v", expEditorialUpdate, actEditorialUpdate)
		}
	}
}

func TestIptcEditor_SetDate(t *testing.T) {
	now := time.Now()
	je := getJpegEditor(LeicaImg, t)
	dates := []IptcDate{DateSent, ReleaseDate, ExpirationDate, DateCreated, DigitalCreationDate}
	for _, d := range dates {
		if err := je.Iptc().SetDate(d, now); err != nil {
			t.Errorf("Could not set date, %v", err)
		}
	}
	//serialise and get Metadata
	md := jpegEditorMD(je, t)
	for _, d := range dates {
		act := md.Iptc().GetDate(d)
		if act.IsZero() {
			t.Errorf("Expected %v data to not be zero", d)
		} else if !cmpDates(act, now) {
			t.Errorf("Expected %v got %v", now, act)
		}
	}

}

func TestIptcEditor_SetEnvelope(t *testing.T) {
	expProductId := []string{"productId1"}
	for _, fname := range []string{LeicaImg, NoExifImg} {
		je := getJpegEditor(fname, t)
		if err := je.Iptc().SetEnvelope(IPTCEnvelope_ProductID, expProductId); err != nil {
			t.Fatalf("Could not set IPTCEnvelope_ProductID for image %s got err %v", fname, err)
		}
		md := jpegEditorMD(je, t)
		actProductId := []string{}
		if err := md.Iptc().ScanEnvelope(IPTCEnvelope_ProductID, &actProductId); err != nil {
			t.Fatalf("Could not scan IPTCEnvelope_ProductId for image %s got err %v", fname, err)
		}
		if !reflect.DeepEqual(actProductId, expProductId) {
			t.Errorf("Expected %v got %v", expProductId, actProductId)
		}
	}
}

func TestIptcEditor_SetKeywords(t *testing.T) {
	expKeywords := []string{"keyword 1", "keyword 2", "keyword 3"}
	for _, fname := range []string{LeicaImg, NoExifImg} {
		je := getJpegEditor(fname, t)

		err := je.Iptc().SetKeywords(expKeywords)
		if err != nil {
			t.Fatalf("Error setting keywords for image %s got error: %v", fname, err)
		}
		md := jpegEditorMD(je, t)
		actKeywords := md.Iptc().GetKeywords()
		if !reflect.DeepEqual(expKeywords, actKeywords) {
			t.Errorf("Expected %v got %v", expKeywords, actKeywords)
		}
	}
}

func TestIptcEditor_SetTag(t *testing.T) {
	expExcursionTolerance := uint8(1)
	for _, fname := range []string{LeicaImg, NoExifImg} {
		je := getJpegEditor(fname, t)
		if err := je.Iptc().Set(IPTCNewsPhoto, IPTCNewsPhoto_ExcursionTolerance, expExcursionTolerance); err != nil {
			t.Fatalf("Could not set IPTCNewsPhoto_ExcursionTolerance for image %s got err %v", fname, err)
		}
		md := jpegEditorMD(je, t)
		actExcursiontolerance := uint8(0)
		if err := md.Iptc().Scan(IPTCNewsPhoto, IPTCNewsPhoto_ExcursionTolerance, &actExcursiontolerance); err != nil {
			t.Fatalf("Could not scan IPTCEnvelope_ProductId for image %s got err %v", fname, err)
		}
		if !(expExcursionTolerance == actExcursiontolerance) {
			t.Errorf("Expected %v got %v", expExcursionTolerance, actExcursiontolerance)
		}
	}

}

func TestIptcEditor_SetTitle(t *testing.T) {
	fnames := []string{LeicaImg, NoExifImg}
	for _, fname := range fnames {
		je := getJpegEditor(fname, t)
		expTitle := "New Title"
		if err := je.Iptc().SetTitle(expTitle); err != nil {
			t.Fatalf("Could not set title for image %s got error %v", fname, err)
		}
		md := jpegEditorMD(je, t)
		if md.Iptc().GetTitle() != expTitle {
			t.Errorf("Expected title %s got %s", expTitle, md.Iptc().GetTitle())
		}
	}

}
