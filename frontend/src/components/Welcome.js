import React from 'react'
import {
    CButton,
    CCard,
    CCardBody,
    CCardHeader,
    CCol,
    CJumbotron,
    CRow,
  } from '@coreui/react'
import { Link } from 'react-router-dom'

function Welcome() {
    return (
        <CRow>
        <CCol>
          <CCard>
            <CCardHeader>
              Welcome
            </CCardHeader>
            <CCardBody>
              <CJumbotron className="border">
                <h1 className="display-3">Selamat Datang</h1>
                <p className="lead">Ini adalah aplikasi sederhana dengan menerapkan <strong>golang</strong> sebagai backend dan <strong>react</strong> sebagai frontend</p>
                <hr className="my-2" />
                <p>Untuk melihat daftar artikel silahkan klik tombol dibawah </p>
                <p className="lead">
                <Link to="/preview">
                  <CButton color="primary" size="lg">Daftar Artikel</CButton>
                  </Link>
                </p>
              </CJumbotron>
            </CCardBody>
          </CCard>
        </CCol>
      </CRow>
    )
}

export default Welcome
