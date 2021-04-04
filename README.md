# pubsub-emulator-toolkit
Tool to better test Google Pub/Sub calls and messages while using the gcloud emulator locally

## Usage
* Start the emulator ([google's docs](https://cloud.google.com/pubsub/docs/emulator))
  ```bash
  $ gcloud components install pubsub-emulator
  $ gcloud components update
  $ gcloud beta emulators pubsub start --project=YOUR_PUBSUB_PROJECTID
* Set the environment variable with the Pub/Sub emulator address
  ```bash
  export PUBSUB_EMULATOR_HOST=localhost:8085
* Run the toolkit
  ```bash
  go run main.go YOUR_PUBSUB_PROJECTID
  
### Available commands
* `newtopic <topic-id>` creates a new topic with name 'topic-id'
* `publish <topic-id> <data> <attributes>` publishes on `topic-id` the message `data` with attributes `attributes`. This is mostly hardcoded, but you can change the message easily in the code itself
* `showtopics` lists all the topics in the active project
* `showsubs` lists all the subscriptions in the active project
* `exit` gracefully exits the program, even if CTRL + C has no side effects.
