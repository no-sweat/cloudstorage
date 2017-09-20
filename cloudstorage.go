package cloudstorage

import (
	"net/http"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"github.com/jung-kurt/gofpdf"
)

func Pdftocloud(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Disposition", "attachment; filename=foo.pdf")
	//w.Header().Set("Content-Type", "application/pdf")

	pdf := gofpdf.New("P", "mm", "A4", "")
			pdf.AddPage()
			pdf.SetFont("Arial", "B", 16)
			pdf.Cell(40, 10, "Hello, world")

	        

	ctx := appengine.NewContext(r)
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	bkt := client.Bucket("no-sweat-pvt.appspot.com")

	if err := bkt.Create(ctx, "no-sweat-pvt", nil); err != nil {
		// TODO: Handle error.
	}

	obj := bkt.Object("data")
	// Write something to obj.
	// w implements io.Writer.
	writer := obj.NewWriter(ctx)
	// Write some text to obj. This will overwrite whatever is there.
	
	/*if _, err := fmt.Fprintf(writer, "This object contains text.\n"); err != nil {
		// TODO: Handle error.
	}*/
	pdf.Output(writer)


	// Close, just like writing a file.
	if err := writer.Close(); err != nil {
		// TODO: Handle error.
	}

	//t := time.Now()
	//http.ServeContent(w, r, "foo.pdf", t, bytes.NewReader(pdf.buffer))

}
