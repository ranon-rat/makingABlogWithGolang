interface Pub {
  Publications: {
    id: number;
    title: string;
    mineatura: string;
    body: string;
  }[];
  Size: number;
}
async function NewPublications() {
  let urlApi: string =
    "api" +
    "/" +
    window.location.pathname.split("/")[window.location.pathname.length - 1];

  let publicationText: string = "";
  let publication: Pub = {
    Publications: [
      {
        id: 0,
        title: "",
        mineatura: "",
        body: "",
      },
    ],
    Size: 0,
  };

  // this is only for get the api
  await fetch(urlApi)
    .then((r) => r.text())
    .then((data) => {
      publicationText = data;
      publication = JSON.parse(publicationText);
    });

  let d: any = document.getElementById("publications");
  for (let i: number = publication.Publications.length - 1; i >= 0; i--) {
    // then add elements into the dom
    let element = `
    <a  class="publications" href="/publication/${publication.Publications[i].id}">
      <div >
        <div class="head">
          <h4 align="center">${publication.Publications[i].title}</h4>
          </div>
            <div class="about" style='background-image:url("${publication.Publications[i].mineatura}")'>
          </div>
      </div>
    </a>
      
    `;
    console.log(publication.Publications[i]);
    d.innerHTML += element;
  }
  let pagePublications: any = document.getElementById("pagePublications");
  for (let i: number = publication.Size / 10; i >= 0; i--) {
    let Element: string = `
 
    <div class="buttonElementID" style="background-color: rgb(255, 255, 255);
     width:1em;
    height: 1em;">
      <a class="buttonElementID" style="background-color: rgb(255, 255, 255);" href="/${Math.round(
        publication.Size / 10 - i
      )}" style=" >
        <div class="buttonElementID" style="  
        
        text-decoration: none;
        color:rgb(0, 0, 0);
        display:inline-block;
        width:10px;
        height: 10px;
        " >
          <h4 > ${Math.round(publication.Size / 10 - i) + 1} </h4>
        </div>
      </a>
    </div >
    `;

    pagePublications.innerHTML += Element;
    console.log(Math.round(publication.Size / 10 - i) + 1);
  }
  console.log(publication);
}
NewPublications();
