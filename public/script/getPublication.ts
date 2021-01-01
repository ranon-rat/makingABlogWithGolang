let urlApi: string =
  "api" +
  "/" +
  window.location.pathname.split("/")[window.location.pathname.length - 1];

let publication: any;
fetch(urlApi)
  .then((r) => r.text())
  .then((data) => {
    publication = data;
  });
console.log(publication);

publication = JSON.parse(publication);
