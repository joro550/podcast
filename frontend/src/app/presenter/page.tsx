import Head from "next/head";
import Image from "next/image";
import "../../../styles/globals.sass";

export default function Page() {
  return (
    <>
      <Head>
        <title>Counter Strike Hot Takes</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="content pt-4">
        <h1>Kassad</h1>
        <hr />
      </main>

      <div className="fixed-grid has-2-cols">
        <div className="grid">

          <div className="cell">
            <div className="card">
              <div className="card-image">
                <figure className="image is-4by3">
                  <Image width={100} height={200} alt={""} src={"/Kassad.jpg"} />
                </figure>
              </div>
              <div className="card-content">
                <div className="media">
                  <div className="media-content">
                    <p className="title is-4">Duncan "Thorin" Shields</p>
                    <p className="subtitle is-6">@Thorin</p>
                  </div>
                </div>
                <p>Lorem ipsum</p>
              </div>
            </div>

          </div>

          <div className="cell">

            <div className="card mb-6">
              <div className="card-header">
                <p className="card-header-title">Episode 1</p>
                <button className="card-header-icon">
                  <span className='icon'>
                    <i className='fas fa-angle-down' aria-hidden="false" ></i>
                  </span>
                </button>
              </div>
              <div className="card-content">
                <table className='table is-striped is-fullwidth'>
                  <thead>
                    <tr>
                      <th>Statement</th>
                      <th>Was Correct?</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>Thing</td>
                      <td>No</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <div className="card mb-6">
              <div className="card-header">
                <p className="card-header-title">Episode 2</p>
                <button className="card-header-icon">
                  <span className='icon'>
                    <i className='fas fa-angle-down' aria-hidden="false" ></i>
                  </span>
                </button>
              </div>
              <div className="card-content">
                <table className='table is-striped is-fullwidth'>
                  <thead>
                    <tr>
                      <th>Statement</th>
                      <th>Was Correct?</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>Thing</td>
                      <td>No</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );

}
