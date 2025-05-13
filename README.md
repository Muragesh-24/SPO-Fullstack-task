# SPO Fullstack Task

This is a backend application built with Go, MongoDB, and hosted on Render. It's designed to collect and store user experiences for the SPO (Student Placement Office) Web Executive Selection process.This project is the second task for the SPO Web Executive selection.

## Live Server

You can view the live version of server: [Live Server](https://spo-fullstack-task.onrender.com) 

## Features
- **Submit Experience**: Allows users to submit their experience with details like company, role, selection process, preparation strategies, and tips.
- **View Total Submissions**: Displays the total number of people who filled out the form.
- **View All Submissions**: Displays all the submissions made by users.

## Project Structure
- **Backend**: Built using Go (Golang) with MongoDB for storing data.
- **Frontend**: A simple HTML, CSS, and JS form for submitting and displaying data.

## How to Run Locally

### Prerequisites
- **Go** (installed on your machine): [Go Install](https://golang.org/doc/install)
- **MongoDB Atlas Account** (for database storage)

### Steps to Run

1. Clone this repository:
   ```bash
   git clone https://github.com/Muragesh-24/SPO-Fullstack-task.git

go mod tidy<br> 
replace mongo connection string in .env file <br> 
go run main.go<br> 
POST /experience: Submit experience form data.<br> 

GET /experience: Get all form submissions.<br> 

GET /experience/stats: Get the total number of form submissions.<br> 