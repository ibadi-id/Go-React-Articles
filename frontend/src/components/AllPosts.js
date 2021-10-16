import React, {useEffect} from 'react'
import {
    CBadge,
    CCard,
    CCardBody,
    CCol,
    CDataTable,
    CRow,
    CNav,
    CNavLink,
    CTabContent,
    CTabPane,
    CTabs,
    CNavItem,
    CButton
  } from '@coreui/react'
import axios from 'axios'
import CIcon from '@coreui/icons-react'
import { Link } from 'react-router-dom'

function AllPosts() {
    const [activeKey, setActiveKey] = React.useState(1)
    
    const [publish, setPublish] = React.useState([])
    const [draft, setDraft] = React.useState([])
    const [trash, setTrash] = React.useState([])
    const fields = ['title','category','action']

    async function fecthArticles(){
        const { data } = await axios.get('http://localhost:3333/article')
        setPublish(data.filter((article)=> article.status === "Publish"))
        setDraft(data.filter((article)=> article.status === "Draft"))
        setTrash(data.filter((article)=> article.status === "Trash"))
    }

    async function trashArticles(item){
        item.status = "Trash"
        const article = {"Article": item}
        await axios.put('http://localhost:3333/article/'+item.id.toString(), article)
        fecthArticles()
    }

    useEffect(() => {
        fecthArticles()
    }, [])
    return (
      <>
      <CTabs>
        <CNav variant="tabs" role="tablist">
          <CNavItem>
            <CNavLink
              href="#"
              active={activeKey === 1}
              onClick={setActiveKey}
              role="tab"
            >
            Published
            <CBadge color="success" className="ml-2 float-right">{publish.length}</CBadge>
            </CNavLink>
          </CNavItem>
          <CNavItem>
            <CNavLink
              href="#"
              active={activeKey === 2}
              onClick={() => setActiveKey(2)}
              role="tab"
            >
              Draft
              <CBadge color="success" className="ml-2 float-right">{draft.length}</CBadge>
            </CNavLink>
          </CNavItem>
          <CNavItem>
            <CNavLink
              href="#"
              active={activeKey === 3}
              onClick={() => setActiveKey(3)}
              role="tab"
            >
            Trashed 
            <CBadge color="success" className="ml-2 float-right">{trash.length}</CBadge>
            </CNavLink>
          </CNavItem>
        </CNav>
        <CTabContent>
          <CTabPane role="tabpanel" aria-labelledby="publish-tab" visible={(activeKey === 1).toString()}>


          <CRow>
          <CCol xs="12" lg="12">
            <CCard>
              <CCardBody>
              <CDataTable
                items={publish}
                fields={fields}
                itemsPerPage={5}
                pagination
                scopedSlots = {{
                  'action':
                  (item)=>(
                    <td>
                      <Link to={`/edit/${item.id}`}>
                        <CButton color="info">
                          <CIcon name="cil-pencil" />
                        </CButton>
                      </Link>
                      <CButton role="tab" aria-selected="true" aria-controls="trash-tab" color="danger" onClick={()=>trashArticles(item)}>
                        <CIcon name="cil-trash" />
                      </CButton>
                    </td>
                  )
                }}
              />
              </CCardBody>
            </CCard>
          </CCol>
          </CRow>



          </CTabPane>

          <CTabPane role="tabpanel" aria-labelledby="draft-tab" visible={(activeKey === 2).toString()}>


            <CRow>
          <CCol xs="12" lg="12">
            <CCard>
              <CCardBody>
              <CDataTable
                items={draft}
                fields={fields}
                itemsPerPage={5}
                pagination
                scopedSlots = {{
                  'action':
                  (item)=>(
                    <td>
                      <Link to={`/edit/${item.id}`}>
                        <CButton color="info">
                          <CIcon name="cil-pencil" />
                        </CButton>
                      </Link>
                      <CButton color="danger" onClick={()=>trashArticles(item)}>
                        <CIcon name="cil-trash" />
                      </CButton>
                    </td>
                  )
                }}
              />
              </CCardBody>
            </CCard>
          </CCol>
          </CRow>


          </CTabPane>

          <CTabPane role="tabpanel" aria-labelledby="trash-tab" visible={(activeKey === 3).toString()}>
            

          <CRow>
          <CCol xs="12" lg="12">
            <CCard>
              <CCardBody>
              <CDataTable
                items={trash}
                fields={fields}
                itemsPerPage={5}
                pagination
                scopedSlots = {{
                  'action':
                  (item)=>(
                    <td>
                      <Link to={`/edit/${item.id}`}>
                        <CButton color="info">
                          <CIcon name="cil-pencil" />
                        </CButton>
                      </Link>
                    </td>
                  )
                }}
              />
              </CCardBody>
            </CCard>
          </CCol>
          </CRow>


          </CTabPane>
        </CTabContent>
      </CTabs>
        
    </>
    )
}

export default AllPosts
