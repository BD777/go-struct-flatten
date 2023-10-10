package pkg

import (
	"testing"

	"github.com/bytedance/mockey"
	"github.com/smartystreets/goconvey/convey"
)

type TestStruct1 struct {
	A string `json:"a"`
	B struct {
		C string `json:"c"`
		D struct {
			E string `json:"e"`
		} `json:"d"`
	} `json:"b"`
	F *string `json:"f"`
	G *struct {
		H string `json:"h"`
	} `json:"g"`
}

func TestStructFlatten(t *testing.T) {

	mockey.PatchConvey("TestStructFlatten", t, func() {

		mockey.PatchConvey("succ", func() {
			obj := TestStruct1{
				A: "a",
				B: struct {
					C string `json:"c"`
					D struct {
						E string `json:"e"`
					} `json:"d"`
				}{
					C: "c",
					D: struct {
						E string `json:"e"`
					}{
						E: "e",
					},
				},
				F: nil,
				G: &struct {
					H string `json:"h"`
				}{
					H: "h",
				},
			}
			flattenData, err := StructFlatten(obj, "json", ".")
			convey.So(err, convey.ShouldBeNil)
			convey.So(flattenData, convey.ShouldNotBeNil)
			convey.So(len(flattenData), convey.ShouldEqual, 5)

			for _, data := range flattenData {
				t.Logf("key: %s, value: %s, kind: %v", data.Key, data.Value.String(), data.Value.Kind())
			}
		})
		mockey.PatchConvey("non struct", func() {
			// non struct
			obj := "string"
			flattenData, err := StructFlatten(obj, "json", ".")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(flattenData, convey.ShouldBeNil)
		})
	})
}
