module.exports = {
  skipFiles: [
    'v0.4/',
    'v0.5/',
    'v0.6/tests',
    'v0.6/vendor',
    'v0.6/VRF*.sol' // tested in Golang
    'v0.7/tests',
    'v0.7/vendor',
    'v0.8/mocks',
    'v0.8/tests',
    'v0.8/vendor',
    'v0.8/VRF*.sol' // tested in Golang
  ],
  mocha: {
    grep: '@skip-coverage', // Find everything with this tag
    invert: true, // Run the grep's inverse set.
  },
}
