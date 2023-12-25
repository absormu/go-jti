
async function getData(){
    console.log('records') 
    const data = await fetch('http://localhost:9670/number-phones');
    const records = await data.json();
    console.log(data) 
    let tab = '';
    records.data.forEach(function(user){
        tab += `<tr>
            <td>${user.id}</td>
            <td>${user.number}</td>
            <td>${user.type}</td> 
            <td>${user.provider.id}</td>
            <td>${user.provider.name}</td>
        </tr>`
    })
    document.getElementById('tbody').innerHTML = tab;
}