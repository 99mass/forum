mobilebtn = document.getElementById("onpen-menu")
mobileMenu = document.getElementById("onclick-mobile-btn")
closebtn = document.getElementById("closebtn")

header = document.getElementById("after-loading")
mainpage = document.getElementById("after-loading1")
loader = document.getElementById("loader")
search = document.getElementById("search")
list_item = document.getElementById("ul-item")



// locations=document.getElementsByClassName("fomat-location")

// locations.forEach(location => {
//     var content=location.value
//     location.Inne = content.split("-").join(" ").split("").join(" ");
// });



setTimeout(() => {
    loader.style.display = "none"
    loader.style.display = "hidden"

    mainpage.style.display = "block"
    mainpage.style.visibility = "visible"

    header.style.display = "block"
    header.style.visibility = "visible"
}, 2000)

mobilebtn.addEventListener("click", () => {
    mobileMenu.style.display = "block"
    mobileMenu.style.visibility = "visible"
    mobileMenu.style.marginLeft = 0
    mobilebtn.style.display = "none"
    mobilebtn.style.display = "hidden"
}, false)

closebtn.addEventListener("click", () => {
    mobilebtn.style.display = "block"
    mobilebtn.style.visibility = "visible"
    // mobileMenu.style.display="none"
    // mobileMenu.style.display="hidden"
    mobileMenu.style.marginLeft = "-100%"
})


locs = document.getElementsByClassName("fomat-location")

// locations.forEach(location => {
//     var content=location.value
//     location.Inne = content.split("-").join(" ").split("").join(" ");
// });
for (var i = 0; i < locs.length; i++) {
    loc = locs[i]
    val = loc.innerHTML
    loc.innerHTML = val.split("-").join(" ").split("_").join(" ");
    console.log(loc.innerHTML)
}
if (search){
    search.addEventListener("keyup", () => {
        var data = search.value
        //    console.log(data)
        var xhr = new XMLHttpRequest();
        xhr.open('GET', '/searchItem?param=' + data, true);
    
        xhr.onload = function () {
            if (xhr.status === 200) {
                var response = JSON.parse(xhr.responseText);
                // console.log(response.Items); 
                if (response.Items.length > 0) {
                    list_item.style.display = "block"
                }
                list_item.innerHTML = ""
                for (let index = 0; index < response.Items.length; index++) {
                    htm_item = `<li class="items" id="item">${response.Items[index]}</li>`
                    list_item.insertAdjacentHTML('beforeend', htm_item)
    
                }
            }
        };
        xhr.send();
    })
}

// search.addEventListener("keydown", () => {
//     elem=document.getElementById("item")
//     console.log(elem)
//     elem.addEventListener("click",()=>{
//         search.value=elem.innerHTML
//     })
//     elem=""
// })
/*execute a function when someone clicks in the document:*/
if(list_item){
    list_item.addEventListener("click", function (e) {
        console.log(e.target.innerHTML)
        search.value = e.target.innerHTML
    });
}
document.addEventListener("click", () => {
    list_item.style.display = "none";
})
// search.addEventListener("mouseout", () => {
//     list_item.style.display="none"
//     }
// )
