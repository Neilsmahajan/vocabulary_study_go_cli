# Vocabulary Study GO CLI

A command-line interface (CLI) tool built with Go to help you study and learn new vocabulary through interactive flashcard sessions.

## Features

- **Interactive Flashcard Sessions**: Learn new words and their definitions in an engaging way.
- **Progress Tracking**: The application remembers which words you know and which you're still learning.
- **Spaced Repetition (Implied)**: By tracking your progress, you can focus on words you don't know yet.
- **Custom Vocabulary**: Easily add your own words to the `vocab.json` file.
- **Track Your Stats**: See how many words you've mastered and how many you have yet to learn.
- **Reset Progress**: Start fresh at any time.

## Installation

1.  **Prerequisites**: Ensure you have [Go](https://golang.org/doc/install) installed on your system.
2.  **Clone the repository**:
    ```bash
    git clone https://github.com/neilsmahajan/vocabulary_study_go_cli.git
    ```
3.  **Navigate to the project directory**:
    ```bash
    cd vocabulary_study_go_cli
    ```
4.  **Build the application**:
    ```bash
    go build -o vocab-cli ./cmd/main.go
    ```

## Usage

### Start a Study Session

To start a flashcard session, simply run the executable:

```bash
./vocab-cli
```

The application will display a word, and you can press `Enter` to reveal the definition. Then, you'll be prompted to mark the word as "known" or "unknown".

### View Statistics

To see your current progress, use the `stats` command:

```bash
./vocab-cli stats
```

This will show you the total number of words, how many you know, how many you don't, and how many you haven't seen yet.

### Reset Progress

To reset all your progress and start over, use the `reset` command:

```bash
./vocab-cli reset
```

You will be asked for confirmation before your progress is erased.

## Configuration

The vocabulary words are stored in the `vocab.json` file. You can edit this file to add, remove, or modify words. The format for each word is as follows:

```json
{
  "word": {
    "part_of_speech": "...",
    "definition": "...",
    "example_sentence": "..."
  }
}
```

## Contributing

Contributions are welcome! If you'd like to improve the application, please follow these steps:

1.  Fork the repository.
2.  Create a new branch (`git checkout -b feature/your-feature-name`).
3.  Make your changes.
4.  Ensure your code is well-formatted using `gofmt`.
5.  Commit your changes (`git commit -m 'Add some feature'`).
6.  Push to the branch (`git push origin feature/your-feature-name`).
7.  Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

- **Name**: Neil Mahajan
- **Email**: [neilsmahajan@gmail.com](mailto:neilsmahajan@gmail.com)
- **Portfolio**: [https://neilsmahajan.com/](https://neilsmahajan.com/)
