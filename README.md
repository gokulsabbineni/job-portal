# job-portal
# Project Overview
Objective: Develop a job portal where users can submit resumes, and the system suggests the best job areas based on their profiles.
Primary Technologies: Golang for backend development, MongoDB for database management.

# System Architecture
Backend: Golang-based RESTful API service.
Database: MongoDB for storing user profiles, job descriptions, and matching criteria.

# Functional Requirements
1.	User account creation and management.
2.	Resume upload and parsing.
3.	Job listing and search functionality.
4.	Automated matching of resumes with job descriptions.

# Database Schema
1.	Users Collection: Stores user information (ID, name, email, etc.).
2.	Resumes Collection: Stores details of resumes (UserID, resume content, parsed data).
3.	Jobs Collection: Stores job listings (JobID, description, requirements).

# API Endpoints (Backend)
1.	User Management:
1.1	POST /users (Register user)
1.2	GET /users/{id} (Retrieve user profile)
1.3	PUT /users/{id} (Update user profile)
2.	Resume Handling:
2.1	POST /resumes (Upload resume)
2.2	GET /resumes/{id} (Retrieve resume details)
3.	Job Management:
3.1	POST /jobs (Create job listing)
3.2	GET /jobs (List jobs)
3.3	GET /jobs/{id} (Job details)
4.	Matching Algorithm:
4.1	POST /match (Match resume with jobs)

