
let updateButton = document.getElementById('updateDetails');
let favDialog = document.getElementById('favDialog');
let outputBox = document.querySelector('output');
let selectEl = document.querySelector('select');
let confirmBtn = document.getElementById('confirmBtn');

const END = 'change';
const START = 'ontouchstart' in document ? 'touchstart' : 'mousedown';
const INPUT = 'input';
const MAX_ROTATION = 35;
const SOFTEN_FACTOR = 3;

function initializeSlimSelect() {
    var my_MSD_object3 = new SlimSelect({
      select: '#contry',
      closeOnSelect: false,
      placeholder: 'Select concerts locations'
    });
}

  // Appeler la fonction pour exécuter le SlimSelect
  initializeSlimSelect();

//Range filter
window.onload = function(){
    slideOne();
    slideTwo();
}

let sliderOne = document.getElementById("slider-1");
let sliderTwo = document.getElementById("slider-2");
let displayValOne = document.getElementById("range1");
let displayValTwo = document.getElementById("range2");
let minGap = 0;
let sliderTrack = document.querySelector(".slider-track");
let sliderMaxValue = document.getElementById("slider-1").max;

function slideOne(){
    if(parseInt(sliderTwo.value) - parseInt(sliderOne.value) <= minGap){
        sliderOne.value = parseInt(sliderTwo.value) - minGap;
    }
    displayValOne.textContent = sliderOne.value;
    fillColor();
}
function slideTwo(){
    if(parseInt(sliderTwo.value) - parseInt(sliderOne.value) <= minGap){
        sliderTwo.value = parseInt(sliderOne.value) + minGap;
    }
    displayValTwo.textContent = sliderTwo.value;
    fillColor();
}
function fillColor(){
    percent1 = (sliderOne.value / sliderMaxValue) * 100;
    percent2 = (sliderTwo.value / sliderMaxValue) * 100;
    sliderTrack.style.background = `#494457`;
}



// Le bouton "Mettre à jour les détails" ouvre le <dialogue> ; modulaire
updateButton.addEventListener('click', function onOpen() {
    if (typeof favDialog.showModal === "function") {
        favDialog.showModal();
    } else {
        console.error("L'API <dialog> n'est pas prise en charge par ce navigateur.");
    }
});
// L'entrée "Animal favori" définit la valeur du bouton d'envoi.
selectEl.addEventListener('change', function onSelect(e) {
    confirmBtn.value = selectEl.value;
});
// Le bouton "Confirmer" du formulaire déclenche la fermeture
// de la boîte de dialogue en raison de [method="dialog"]
favDialog.addEventListener('close', function onClose() {
    outputBox.value = favDialog.returnValue + " bouton cliqué - " + (new Date()).toString();
});


$(document).on("pagecreate", function () {
    $("#foo").on("click", function () {
        // close button
        var closeBtn = $('<a href="#" data-rel="back" class="ui-btn-right ui-btn ui-btn-b ui-corner-all ui-btn-icon-notext ui-icon-delete ui-shadow">Close</a>');

        // text you get from Ajax
        var content = "<p>Lorem ipsum dolor sit amet, consectetur adipiscing. Morbi convallis sem et dui sollicitudin tincidunt.</p>";

        // Popup body - set width is optional - append button and Ajax msg
        var popup = $("<div/>", {
            "data-role": "popup"
        }).css({
            width: $(window).width() / 2.5 + "px",
            padding: 5 + "px"
        }).append(closeBtn).append(content);

        // Append it to active page
        $.mobile.pageContainer.append(popup);

        // Create it and add listener to delete it once it's closed
        // open it
        $("[data-role=popup]").popup({
            dismissible: false,
            history: false,
            theme: "b",
            /* or a */
            positionTo: "window",
            overlayTheme: "b",
            /* "b" is recommended for overlay */
            transition: "pop",
            beforeposition: function () {
                $.mobile.pageContainer.pagecontainer("getActivePage")
                    .addClass("blur-filter");
            },
            afterclose: function () {
                $(this).remove();
                $(".blur-filter").removeClass("blur-filter");
            },
            afteropen: function () {
                /* do something */
            }
        }).popup("open");
    });
});
