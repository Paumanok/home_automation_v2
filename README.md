# Home_Automation_v2
### Here we go again

[Basic API documentation](internal/server/readme.md)

## Purpose: 
To recreate my first attempt at a home automation web app in a more performant and useful form. 

The first version worked, great. 2ish files of python flask and some mongodb. I learned some html templating, some new python concepts. 
But it was slow. To retrieve the last month or so of measurements through mongodb would take nearly 30 seconds of screaming unoptimized nosql. 

I've been curious about Go for a while so I figured I'd take a more serious stab at it. 

## Components
- Golang http server, api routing by hand, pgx & squirrel for the db
- Postgresql to learn some sql
- One nice docker bow around it. 
- My [ESP32 environment sensing project](https://github.com/Paumanok/esp_environment_sensing)
- (For now) My [old server](https://github.com/Paumanok/home_automation) for forwarding data while I develop


## Future goals
- Learn a frontend framework enough to make it pretty
- golang Telegram bot integration for data retrieval and other tasks


## Thanks
- [bnkamalesh's goapp project structure](https://github.com/bnkamalesh/goapp) for giving me an excellent visual reference for structuring everything.
