import { extend } from 'vee-validate'
import { required, size } from 'vee-validate/dist/rules'

extend('required', {
  ...required,
  message: 'This field is required',
})
extend('size', {
  ...size,
  message: 'File is larger than the configured file size limit.',
})
extend('ext', {
  ...size,
  message: 'File type is not allowed.',
})
