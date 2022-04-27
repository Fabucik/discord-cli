# discord-cli
(IN DEVELOPMENT)

---

## LINUX ONLY FOR NOW

CLI tool to get data about your discord servers and profile

**Go language is required, if you don't have it installed, see: https://go.dev/doc/install**

# Get Started
- Create file called ".env" in this project's directory
- Turn on developer mode in your Discord account
- Go to https://discord.com/developers/applications
- Create new application
- Copy Client Secret and Client ID, and put them in .env file in this format: CLIENTID=<id>
                                                                              CLIENTSECRET=<secret>
- cd to project's directory and run: make
- Add /your/directory/discord-cli/bin to the PATH environment variable using `export PATH=$PATH:/your/directory/discord-cli/bin` (replace /your/directory with directory where discord-cli is installed)