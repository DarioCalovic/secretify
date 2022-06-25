<template>
  <div>
    <section class="section is-medium">
      <div class="container darkness">
        <div class="columns is-centered">
          <div class="column is-12">
            <h1 class="title">Imprint</h1>
            <h2 class="title is-4">Kontakt-Addresse</h2>
            <p>
              Dario Calovic<br />Hobacherweg 13<br />6010 Kriens<br />Schweiz
            </p>
            <p>E-Mail:<br />dario.calovic@gmail.com</p>
            <h2 class="title is-4">Haftungsausschluss</h2>
            <p>
              Der Autor übernimmt keinerlei Gewähr hinsichtlich der inhaltlichen
              Richtigkeit, Genauigkeit, Aktualität, Zuverlässigkeit und
              Vollständigkeit der Informationen.
            </p>
            <p>
              Haftungsansprüche gegen den Autor wegen Schäden materieller oder
              immaterieller Art, welche aus dem Zugriff oder der Nutzung bzw.
              Nichtnutzung der veröffentlichten Informationen, durch Missbrauch
              der Verbindung oder durch technische Störungen entstanden sind,
              werden ausgeschlossen.
            </p>
            <p>
              Alle Angebote sind unverbindlich. Der Autor behält es sich
              ausdrücklich vor, Teile der Seiten oder das gesamte Angebot ohne
              besondere Ankündigung zu verändern, zu ergänzen, zu löschen oder
              die Veröffentlichung zeitweise oder endgültig einzustellen.
            </p>
            <h2 class="title is-4">Haftungsausschluss für Links</h2>
            <p>
              Verweise und Links auf Webseiten Dritter liegen ausserhalb unseres
              Verantwortungsbereichs. Es wird jegliche Verantwortung für solche
              Webseiten abgelehnt. Der Zugriff und die Nutzung solcher Webseiten
              erfolgen auf eigene Gefahr des jeweiligen Nutzers.
            </p>
            <h2 class="title is-4">Urheberrechte</h2>
            <p>
              Die Urheber- und alle anderen Rechte an Inhalten, Bildern, Fotos
              oder anderen Dateien auf dieser Website, gehören ausschliesslich
              <strong>Dario Calovic</strong> oder den speziell genannten
              Rechteinhabern. Für die Reproduktion jeglicher Elemente ist die
              schriftliche Zustimmung des Urheberrechtsträgers im Voraus
              einzuholen.
            </p>
            <p>
              Source (Self-translation):
              <a
                href="https://www.swissanwalt.ch"
                target="_blank"
                rel="noopener"
                >SwissAnwalt</a
              >
            </p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
export default {
  name: 'Imprint',
  mounted() {
    // Track
    this.$track.pageview({})
  },
  head() {
    return {
      title: 'Imprint - Secretify',
    }
  },
}
</script>
