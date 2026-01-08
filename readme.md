# ASCII Art Web

## Description

ASCII Art Web is a Go-based web application that converts user-provided text into ASCII art using different banner styles (Standard, Shadow, Thinkertoy). The application validates input, applies the selected ASCII font, and renders the result directly in the browser while preserving formatting such as line breaks.

This project demonstrates the use of HTTP handlers in Go, form processing, input validation, template rendering, and string manipulation for text-based visual output.

---

## Authors

* Oumaima Talhaoui          (otalhaou)
* Baha-eddine Aboulouafa    (baboulou)
* Kaoutar Hammaoui          (khammaou)

---

## Usage: how to run

### Steps

1. Clone the repository:

   ```bash
   git clone https://learn.zone01oujda.ma/git/otalhaou/ascii-art-web
   cd ascii-art-web
   ```

2. Run the server:

   ```bash
   go run .
   ```

3. Open your browser and navigate to:

   ```text
   http://localhost:8080
   ```

4. Enter text in the textarea, choose an ASCII art style, and click **Afficher** to see the generated ASCII art.

---

## Implementation details: algorithm

1. **HTTP Request Handling**

   * The `/` endpoint only accepts `GET` requests.
   * The `/ascii-art` endpoint only accepts `POST` requests.
   * Any other HTTP method returns `405 Method Not Allowed` status.

2. **Input Parsing and Validation**

   * The input text is read from the form field `content`.
   * The program checks:

     * The input is not empty.
     * All characters are valid ASCII characters (range 32–126, plus newline characters).
     * The input length does not exceed 3000 characters.

3. **Font Selection**

   * The selected ASCII art font is read from the `types` radio button.
   * If no font is selected, an error message is returned.

4. **ASCII Art Generation**

   * The validated input text and selected font are passed to the `ApplyingFont` function.
   * This function maps each character to its ASCII art representation using the chosen banner file.

5. **Newline Handling**

   * If the generated ASCII art starts with a newline, an extra newline is added to both the result and the original input.
   * This ensures consistent rendering between the `<pre>` block (ASCII result) and the `<textarea>` (user input).

6. **Template Rendering**

   * The `home.html` template is parsed and executed.
   * The page displays:

     * The original input text
     * The selected font
     * The generated ASCII art
     * Any validation or processing errors

This structured flow ensures clean separation between validation, processing, and presentation while handling edge cases related to formatting and user input.
