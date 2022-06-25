import SettingsRepository from '~/repositories/SettingsRepository'
import SecretRepository from '~/repositories/SecretRepository'
import FileRepository from '~/repositories/FileRepository'

export default ($axios, apiURL) => ({
  settings: SettingsRepository($axios, apiURL),
  secret: SecretRepository($axios, apiURL),
  file: FileRepository($axios, apiURL),
})
