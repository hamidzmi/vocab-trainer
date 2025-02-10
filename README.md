# German Vocabulary Trainer (CLI) 🎓🇩🇪

A **Command-Line Application** built in **Go** to help users learn and practice German vocabulary using an interactive menu-driven approach. The app supports adding new words, reviewing learned vocabulary, and testing knowledge through quizzes.

## Features
👉 **AI-Powered Translations & Examples** – Uses Generative AI (via Ollama) to fetch word meanings and create contextual example sentences.  
👉 **Learn New Words** – Learn new words with AI-powered translations and examples.  
👉 **Review Mode** – Go through previously learned words.  
👉 **Quiz Mode** – Test your vocabulary knowledge interactively.  

## Prerequisites

Before running the application, ensure that **Ollama** is installed and running:

### Install Ollama

Follow the installation guide on the [Ollama website](https://ollama.ai/) to set up Ollama on your system.

### Start Ollama

Ensure that Ollama is running before starting the vocabulary trainer:

```sh
ollama serve
```

## Installation

```sh
git clone https://github.com/yourusername/german-vocab-trainer.git
cd german-vocab-trainer
go build -o vocab-trainer
./vocab-trainer
```

## Usage

Run the program and choose an option:

```
Let me know what do you want to do now:
1. Learn a new word
2. Review learnt words
3. Quiz
4. Exit
Enter your choice:
```

## Running Tests

```sh
go test ./...
```

## Roadmap 🛤️
🔹 Integrate AI-generated translations and example sentences (via Ollama)  
🔹 Implement Spaced Repetition for better retention  
🔹 Implement Review mode  
🔹 Implement Quiz mode  
🔹 Expand to GUI/iOS app  

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License
MIT License
