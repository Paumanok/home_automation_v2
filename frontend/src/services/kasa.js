
export async function toggle_lamp() {
  
  const response = fetch("http://kasa.datapaddock.lan/toggle", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    },
    body: '{"idx": 0}'
  });
  

  //response.then((resp) => {
  //  const json_promise = resp.json();

  //  json_promise.then((data) => {
  //    next.value = data["sync_time"]
  //  })
  //});
}
