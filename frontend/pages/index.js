import React from 'react'
import fetch from 'isomorphic-fetch'

export default class Index extends React.Component {
  static async getInitialProps () {
    console.log('hihi')
  }
  render () {
    return (
      <div>hihi</div>
    )
  }
}
