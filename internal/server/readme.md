# Unofficial API Documentation

## Putting this here so I don't forget how to use my own tool

# /api
--- 
## ../measurements

### ../last

#### Query Params
- period -- Period of time to return recent measurements, most recent being the last value
    - last -- Only the last measurement recieved per device
    - hour
    - day 
    - week 
    - month
- comp
    - Return measurements with pre-calculated 
    - default: false
- byDevice
    - Pre-sort measurements by device with added device metadata
    - default false
example:
` api/measurements/last?period=hour&comp=true&byDevice=true `

### ../range

#### Query Params
- start -- start of range
    - url encoded datetime 

- end -- end of range
    - url encoded datetime 

example: 
` api/measurements/range?end=2023-06-16T01%3A04%3A46.814883%2B00%3A00&start=2023-06-15T01%3A04%3A46.814883%2B00%3A00`


## ../devices

### default -- api/devices
    - retuns all registered devices in db

### ../update

#### Query Params
- mac --string
    - required, tells update which device to modify
- nickname -- string
    - update nickname
- humidity_comp -- int
    - update humidity compensation
- temp_comp -- int
    - update temperature compensation

example:
`"api/devices/update?mac=84:f7:03:f1:1b:62&hum_comp=2&nickname=lovelyday&temp_comp=-2`  
