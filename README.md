# Groupie Tracker

Welcome to **Groupie Tracker**! This project lets you dive into the vibrant world of music by harnessing the power of an API to bring you closer to your favorite artists, concerts, and locations. Let’s turn data into a symphony of information!

## Overview

Groupie Tracker is designed to make the most out of a comprehensive music API, breaking it down into four major components:

- **Artists**: Discover your favorite bands and solo acts! Get insights into their names, images, the years they kicked off their careers, the release dates of their first albums, and the members who make up the magic.
  
- **Locations**: Stay updated on where the latest concerts are happening. Never miss a beat with information about past and upcoming events.

- **Dates**: Find out when the music is playing! Access a complete list of concert dates that you can’t afford to miss.

- **Relations**: The glue that holds it all together! This part connects artists, dates, and locations, creating a seamless experience for music lovers.

In addition, we focus on creating and visualizing events and actions, utilizing a robust client-server architecture to enhance user experience.

### Project Structure

Here's a quick glance at the structure of the project:

```
/home/hilaromondi/groupie-tracker/
├── api/
│   ├── fetchartist.go
│   ├── fetchdates.go
│   ├── fetchlocations.go
│   └── fetchrelations.go
├── handlers/
│   ├── artistdetails.go
│   ├── dateshandler.go
│   ├── HandleError.go
│   ├── homehandler.go
│   ├── locationhandler.go
│   └── relationhandler.go
├── models/models.go
├── static/
│   ├── 400.css
│   ├── 404.css
│   ├── 500.css
│   ├── main.css
│   └── ... (other static files)
├── templates/
│   ├── 400.html
│   ├── 404.html
│   ├── 500.html
│   ├── index.html
│   └── ... (other templates)
├── go.mod
├── guide_txt.md
├── main.go
└── README.md
```

## Installation

Ready to rock? First, ensure you have the latest version of Go installed on your machine. If you don’t have it yet, grab it from the official [Go website](https://golang.org/dl/).

To clone this project to your local machine, use the following command:

```bash
git clone https://learn.zone01kisumu.ke/git/hilaromondi/groupie-tracker.git
```

## Usage

Once you’ve cloned the repository, navigate into the project folder:

```bash
cd groupie-tracker
```

Then, fire it up with:

```bash
go run main.go
```

Open your browser and paste the following link to get started: `http://localhost:8082`. Let the music play!

## Contributing

We love collaboration! Pull requests are welcome, and for major changes, please open an issue first to discuss your ideas. Let’s make this project even better together!

### Contributors

- [bobaigwa](https://learn.zone01kisumu.ke/git/bobaigwa)
- [hilaromondi](https://learn.zone01kisumu.ke/git/hilaromondi/groupie-tracker)
- [svictor](https://learn.zone01kisumu.ke/git/svictor)

## License

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/). Rock on! 🎸