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
    <div >
      <div class="head">
      <h2 align="center">${i.title}</h2>
      </div>
      <div class="about" style='background-image:url("${i.mineatura}")'>
      
    </div>
      </div>
      </a>
    `;

    d.innerHTML += element;
  }
  let pagePublications = document.getElementById("pagePublications");
  for (let i: number = 0; i <= publication.Size / 10; i++) {
    let Element: string = `
    <a class="buttonElementID" href="/${i}">
      <div >
        <p> ${i} </p>
      <div>
    </a>
    `;
    pagePublications.innerHTML += Element;
    console.log(i);
  }
  console.log(publication);
}
NewPublications();
