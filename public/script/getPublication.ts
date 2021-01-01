interface infoDocument {
  id: number;
  title: string;
  mineatura: string;
  bodyOfDocument: string;
}
interface Api {
  Publications: infoDocument[];
}

async function NewPublications() {
  let urlApi: string =
    "api" +
    "/" +
    window.location.pathname.split("/")[window.location.pathname.length - 1];

  let publication: any;
  // this is only for get the api
  await fetch(urlApi)
    .then((r) => r.text())
    .then((data) => {
      publication = data;
    });

  publication = await JSON.parse(publication);

  let d: HTMLElement = document.getElementById("publications");
  for (let i of publication.Publications) {
    // then add elements into the dom
    let element = `
    <a  class="publications" href="/publication/${i.id}">
    <div class="publications">
    
      <h2>${i.title}</h2>
      <img height =300 src= "${i.mineatura}">
    
      </div>
      </a>
    `;

    d.innerHTML += element;
  }
}
NewPublications();
