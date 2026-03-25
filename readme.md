# PDF Transformers

**PDF Transformers** is a Go program that converts a PDF file into PNG images and generates a JSON file listing all created images. Each page of the PDF becomes a separate PNG image.  

---

## Features

- Convert a PDF file in a specified folder to individual PNG images (one per page)  
- Generate a JSON file with a list of all created PNG files  
- Output files are stored in a `slike` folder inside the user's folder  

---

## Setup Instructions

1. **Clone the repository**  

   Place the repository on your local machine, for example:  
   `D:\go_workspace\pdf-transformers`  

2. **Install Poppler**  

   - Download Poppler for Windows from the internet  
   - Extract it to `D:\poppler`  
   - Ensure `pdftoppm.exe` is located at `D:\poppler\Library\bin\pdftoppm.exe`  

3. **Update the code for local testing**  

   - In `main.go`, uncomment the following two lines:  
     ```go
     basePath := "D:\\go_workspace\\pdf-transformers"
     "D:\\poppler\\Library\\bin\\pdftoppm.exe"
     ```  
   - This sets the local paths for testing  

---

## Running the Program

### 1. Directly with Go

Navigate to the folder `pdf-transformers` and run:

```bash
go run main.go <folder_name> <pdf_file>