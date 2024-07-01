
export async function get_next(next) {
  
  const response = fetch("/api/next", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    },
  });
  

  response.then((resp) => {
    const json_promise = resp.json();

    json_promise.then((data) => {
      next.value = data["sync_time"]
    })
  });
  //const resp_json = await response.json()

  //resp_json.then( (data) => {
  //  next.value = data["sync_time"]
  //})
  //next.value = resp_json["sync_time"]
}


export async function get_measurements() {
  const response = await  fetch("/api/measurements/last?period=hour&byDevice=true&comp=true&fahrenheit=true", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    },
  });
  
  const resp_json = await response.json()
  return resp_json
  //console.log(resp_json)
 // response.then((resp) => {
 //     const json_promise = resp.json();

 //     json_promise.then((data) => {
 //       return data
 //     })
 //   });
 // return null
}


