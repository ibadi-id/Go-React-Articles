import React, { useEffect } from 'react'
import {
  CCard,
  CCardBody,
  CCardHeader,
  CCol,
  CRow,
  CCollapse,
  CFade,
  CBadge,
  CPagination
} from  '@coreui/react'
import axios from 'axios'

function Preview() {
    const collapsed = true
    const showCard = true
    const [articles, setArticles] = React.useState([])
    const [currentPage, setCurrentPage] = React.useState(1)
    const [pageNumber, setPageNumber] = React.useState(1)

    async function fecthArticles(offset){
        if (offset == null) {
            const { data } = await axios.get('http://localhost:3333/article')
            const publish = data.filter((article)=> article.status === "Publish")
            setArticles(publish.slice(0, 6))
            setPageNumber(Math.ceil(publish.length/6))
        } else {
            const { data } = await axios.get('http://localhost:3333/article/6/'+offset)
            setArticles(data.filter((article)=> article.status === "Publish"))
        }
    }

    useEffect(() => {
        fecthArticles()
    }, [])

    function nextPage(e){
        let offset = (e-1)*6
        fecthArticles(offset.toString())
        setCurrentPage(e)
    }
    return (
        <>
        <CRow>
            {articles.map((article, id) =>(
                <CCol key={id} xs="12" sm="6" md="4">
                <CFade  in={showCard}>
                    <CCard>
                    <CCardHeader>
                        {article.title}
                        <div className="card-header-actions">
                            <CBadge color="success" className="float-right">{article.status}</CBadge>
                        </div>
                    </CCardHeader>
                    <CCollapse show={collapsed}>
                        <CCardBody>
                        {article.content}
                        </CCardBody>
                    </CCollapse>
                    </CCard>
                </CFade>
                </CCol>
             ))}
        </CRow>
        <CPagination
            size="lg"
            align="end"
            activePage={currentPage}
            pages={pageNumber}
            onActivePageChange={(e)=>nextPage(e)}
          />
        </>
    )
}

export default Preview
