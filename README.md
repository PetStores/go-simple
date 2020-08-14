# go-simple
A simple HTTP-based Petstore backend written in Go. 

## Limitations

### Configuration

#### Secret management

We don't use any secret management practices in this application.
You can see that the database connection string is provided via the `DATABASE_URL`
environment variable.

If you build a highly secure application, please consider using a proper solution instead.
For instance, you might be interested in Vault.
