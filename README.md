This is a proof of concept for a bot that notifies a Slack channel when a new release for a specific repo (this one) has been created on GitHub. To work, it requires a webhook for releases to be set up on GitHub. The webhook hits the `/githubwebhook` route for the server with a JSON payload describing the release. Upon receipt of this payload, the server sends a notification to a dummy Slack workspace set up for this purpose. 

#How To Use (in a development environment)#
1. Install [ngrok](https://ngrok.com/).
2. Start ngrok: `ngrok http 8080`. The app is set to listen on port 8080.
3. Create a new webhook for your GitHub repository (Settings > Webhooks > Add webhook). 
* Use the ngrok URL (displayed in the terminal when you started it) as the payload URL.
* Change the content type to "application/json".
* Select "Let me select individual events" and pick the Release event. 
* Check the checkbox next to "Active".
* Add the webhook.
4. Add to Slack.
* Create a new incoming webhook for your workspace [here](https://my.slack.com/services/new/incoming-webhook/)
* Save the webhook URL. It will be necessary for the next step.
5. In a different terminal, create an environment variable named `SLACK_URL` and set it to the wehbook URL. Without this environment variable, the app will fail.
6. In the same terminal window, build and run the app like normal. Whenever a release is created, the Slack channel(s) you chose will get a message with the version number.
