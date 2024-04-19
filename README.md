**To whoever is reviewing this code:**

There are a few things I want to explain

1. In the unit tests of this project, I forgot how to mock the DB behavior, resulting in "invalid address" issues for the parts involving DB operations in the tests.

2. I'm not very familiar with using Docker, but I still tried to write a Dockerfile and docker-compose.yml. However, I couldn't find a way to connect the Go service with the MySQL container in Docker. So, I ended up running the program locally.

3. I wanted to explain that although this challenge was sent on Monday night, I actually started working on this project only on Wednesday afternoon due to presentations I had to prepare for my school courses. The time was a bit tight for me. There are many areas that are not optimized, so please understand.

Thank you for your understanding.

***Usage of API key*** 

Just set the value of API_KEY to a valid value in the **common/const** directory

Now, you can run my code by **sh ./start.sh**