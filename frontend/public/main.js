const button = document.getElementById("bt");
const textBox = document.getElementById("pData");

button.addEventListener("click", getBackendValues);

function getBackendValues(event){
    fetch('/api/ab',
    {
        method: 'GET'
    })
    .then((response) => {
        if(response.ok){
            return response.json()
        } else {
            console.log(response)
        }
    })
    .then((json) => {
        textBox.innerText = JSON.stringify(json);
    })
}