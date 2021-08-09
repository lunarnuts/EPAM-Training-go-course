### HomeWork08
1. Take HomeWork 6 Task 3 as initial step.
2. Create new service to store user-tokens from Web-form(no need to use 
databases. Feel free to use map to store it).
3. Refactor existing POST-request store tokens on new service. New service 
should store them as structure{token: your token, createdAt: currentdate, 
expiredAt: currentdate + 10 days}.
4. Make DockerFiles to both services
5. Create account on docker hub. 
6. Create images, tag it, upload it on your own hub. 
7. Make Docker Compose File
8. Commit DockerFiles and Docker Compose File in Git, open PR. 
9. Share link with your mentors on PR and images in your own docker hub. 
Way to success: Your app should works correct in docker container.
