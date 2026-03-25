PDF Transformers

PDF Transformers is a Go program that converts a PDF file into PNG images and generates a JSON file listing all created images. Each page of the PDF becomes a separate PNG image.

Features
Convert a PDF file in a specified folder to individual PNG images (one per page)
Generate a JSON file with a list of all created PNG files
Output files are stored in a slike folder inside the user's folder
Setup Instructions

Clone the repository:

Place the repository on your local machine, for example:
D:\go_workspace\pdf-transformers

Install Poppler:
Download Poppler for Windows from the internet
Extract it to D:\poppler
Ensure pdftoppm.exe is located at D:\poppler\Library\bin\pdftoppm.exe
Update the code for local testing:

In main.go, uncomment the following two lines:

basePath := "D:\\go_workspace\\pdf-transformers"
"D:\\poppler\\Library\\bin\\pdftoppm.exe"
This sets the local paths for testing
Running the Program
1. Directly with Go

Navigate to the folder pdf-transformers and run:

go run main.go <folder_name> <pdf_file>

Example:

go run main.go RDC_BGD primer.pdf

This will:

Convert primer.pdf located in the folder RDC_BGD
Create PNG images for each page inside D:\go_workspace\pdf-transformers\RDC_BGD\slike
Generate a JSON file primer.json listing all PNG files
2. Build an EXE

To create a standalone executable:

go build -o pdf_transformers.exe main.go

Then you can run it like this:

.\pdf_transformers.exe RDC_BGD primer.pdf
Folder Structure Example

After running the program, your folder structure will look like this:

D:\go_workspace\pdf-transformers\
├─ main.go
├─ pdf_transformers.exe
├─ RDC_BGD\
│   ├─ primer.pdf
│   └─ slike\
│       ├─ primer_strana-1.png
│       ├─ primer_strana-2.png
│       ├─ primer_strana-3.png
│       └─ primer.json
Example JSON Output (primer.json)
[
  "primer_strana-1.png",
  "primer_strana-2.png",
  "primer_strana-3.png"
]
The JSON file contains only the names of the generated PNG files
Useful for automatic processing or displaying the images in applications
Notes
Ensure that Poppler is correctly installed in D:\poppler
The slike folder is created automatically if it does not exist
JSON file will be named after the PDF (without .pdf)
Each page of the PDF corresponds to a PNG file named: <PDFName>_strana-<page_number>.png