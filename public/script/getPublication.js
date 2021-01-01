var urlApi = "api" +
    "/" +
    window.location.pathname.split("/")[window.location.pathname.length - 1];
var publication;
fetch(urlApi)
    .then(function (r) { return r.text(); })
    .then(function (data) {
    publication = data;
});
console.log(publication);
publication = JSON.parse(publication);
