import io
from PyPDF2 import PdfReader
import re
import sys

def extract_technical_skills_from_pdf_content(pdf_content):
    reader = PdfReader(io.BytesIO(pdf_content))
    full_text = ""
    for page_number, page in enumerate(reader.pages, 1):
        page_text = page.extract_text()
        full_text += page_text + "\n"

    start_pattern = "Technical  Skills"
    end_pattern = "PROFESSIONAL  EXPERIENCE"

    skills_match = re.search(f"{start_pattern}(.*?){end_pattern}", full_text, re.DOTALL)

    if skills_match:
        extracted_text = skills_match.group(1).strip()
        return extracted_text
    else:
        return "Technical skills section not found."

if __name__ == "__main__":
    pdf_content = sys.stdin.buffer.read()
    skills_text = extract_technical_skills_from_pdf_content(pdf_content)
    print(skills_text)
