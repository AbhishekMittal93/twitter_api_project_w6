# Twitter API Integration with Golang

This project demonstrates how to interact with the Twitter API using Golang. It has two main functionalities: Posting a new tweet and Deleting an existing tweet. The project uses OAuth 1.0a for user authentication and allows interaction with Twitter’s API endpoints to manage tweets through programs.

## Functionality

- **Post a New Tweet**: Post a message to your Twitter account using the Twitter API.
- **Delete a Tweet**: Delete an existing tweet by providing its tweet ID.

## Prerequisites
- A Twitter Developer account with API keys and tokens generated.
- Create a .env file in the root directory of the project with the following keys:
```env
TWITTER_API_KEY=your-api-key
TWITTER_API_SECRET_KEY=your-api-secret-key
TWITTER_ACCESS_TOKEN=your-access-token
TWITTER_ACCESS_TOKEN_SECRET=your-access-token-secret
```
- A valid .env file containing the Twitter API keys and tokens.

## Technologies Used

- **Golang**: Programming language used for building the application.
- **Twitter API**: The external API used to interact with Twitter.
- **OAuth 1.0a**: Authentication method used to securely interact with the Twitter API.
