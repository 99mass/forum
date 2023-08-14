
map_cadr = document.getElementById("map-card")

// Clé API MapQuest
const apiKey = '2Q1G1qJUEd1zfcWDDaWs9klqdft8tfvN';

// Initialiser la carte
L.mapquest.key = apiKey;

// Options de la carte
const mapOptions = {
    center: [48.8566, 2.3522], // Coordonnées de Paris (par exemple)
    layers: L.mapquest.tileLayer('map'),
    zoom: 12
};

// Créer la carte dans le conteneur spécifié (utilisation de l'ID "map-card")
const map = L.mapquest.map(map_cadr, mapOptions);

// Liste de coordonnées des points d'intérêt
const pointsOfInterest = [
    { lat: 48.8566, lng: 2.3522 }, // Paris
    { lat: 40.7128, lng: -74.0060 }, // New York
    { lat: 51.5074, lng: -0.1278 }, // Londres
    // Ajoutez d'autres coordonnées ici
];

// Icône personnalisée pour les marqueurs
// const customIcon = L.icon({
//     iconUrl: 'chemin/vers/votre/icon.png',
//     iconSize: [32, 32], // Taille de l'icône en pixels
//     iconAnchor: [16, 32], // Point d'ancrage de l'icône en pixels (position de pointe)
// });

// Ajouter les marqueurs à la carte
pointsOfInterest.forEach(point => {
    L.marker([point.lat, point.lng], { icon: customIcon }).addTo(map);
});
