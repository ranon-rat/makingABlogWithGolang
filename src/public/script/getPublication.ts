interface Pub {
  Cantidad: number;
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

  let publication: Pub | string = {
    Publications: [
      {
        id: 0,
        title: "",
        mineatura: "",
        body: "",
      },
    ],
    Size: 0,
    Cantidad: 0,
  };

  // this is only for get the api
  await fetch(urlApi)
    .then((r) => r.text())
    .then((data) => {
      console.log(data);
      publication = data;
      publication = JSON.parse(publication);
    });

  let d: any = document.getElementById("publications");
  for (let i of publication.Publications) {
    // then add elements into the dom
    let element = `
    <p>
    <a  class="publications" href="/publication/${i.id}">
      <div >
        <div class="publicationContent">
        <img src="${i.mineatura}"style="height: 8em;" > 
              
        <h4 >
            ${i.title}
          </h4>
         
          
          </div>
        </div>
      </div>
    </a>
    </p>
      
    `;

    d.innerHTML += element;
  }
  let pagePublications: any = document.getElementById("pagePublications");
  for (let i: number = publication.Size / publication.Cantidad; i > 0; i--) {
    let Element: string = `
 
    <div class="buttonElementID" style="background-color: rgb(255, 255, 255);
     width:1em;
    height: 1em;">
      <a class="buttonElementID" style="background-color: rgb(255, 255, 255);" href="/${Math.round(
        publication.Size / publication.Cantidad - i
      )}" style=" >
        <div class="buttonElementID" style="  
        
        text-decoration: none;
        color:rgb(0, 0, 0);
        display:inline-block;
        width:10px;
        height: 10px;
        " >
          <h4 > ${
            Math.round(publication.Size / publication.Cantidad - i) + 1
          } </h4>
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
