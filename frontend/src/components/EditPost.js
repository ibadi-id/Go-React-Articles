import React, {useEffect} from 'react'
import {
    CButton,
    CCard,
    CCardBody,
    CCardHeader,
    CCol,
    CForm,
    CFormGroup,
    CTextarea,
    CInput,
    CLabel,
    CRow,
  } from '@coreui/react'
import { useHistory } from "react-router-dom";
import axios from 'axios'


function EditPost({match}) {
    let history = useHistory();
    const [validated, setValidated] = React.useState(false)
    const [title, setTitle] = React.useState("")
    const [content, setContent] = React.useState("")
    const [category, setCategory] = React.useState("")
    const [status, setStatus] = React.useState("Draft")

    useEffect(() => {
        async function fecthArticle(){
            const { data } = await axios.get(`http://localhost:3333/article/${match.params.id}`)
            setTitle(data.title)
            setContent(data.content)
            setCategory(data.category)
            setStatus(data.status)
        }
        fecthArticle()
    }, [match])

    

    async function updateArticles(item){
        const article = {"Article": item}
        await axios.put(`http://localhost:3333/article/${match.params.id}`, article)
    }

    async function handleSubmit(event){
        event.preventDefault()
        const form = event.currentTarget
        if (form.checkValidity() === false) {
            event.preventDefault()
            event.stopPropagation()
            setValidated(false)
        } else {
            setValidated(true)
            const item = {}
            item.title = title
            item.content = content
            item.category = category 
            item.status = status
            
            await updateArticles(item)
            history.push("/preview");
        }
    }
    return (
        <CRow>
            <CCol xs="12" sm="6">
          <CCard>
            <CCardHeader>
                Edit Post
            </CCardHeader>
            <CCardBody>
            <CForm className="was-validated"
                action="" method="post"  
                validated={validated.toString()}
                onSubmit={handleSubmit} noValidate>
              <CFormGroup>
                <CLabel htmlFor="title">Title</CLabel>
                <CInput id="title" placeholder="Input Title" minLength="20" maxLength="200" value={title} onChange={(e)=>setTitle(e.target.value)} required/>
              </CFormGroup>
              <CFormGroup>
                <CLabel htmlFor="content">Content</CLabel>
                    <CTextarea 
                      name="textarea-input" 
                      id="content" 
                      rows="9"
                      type='text'
                      placeholder="Input Content..." 
                      value={content} onChange={(e)=>setContent(e.target.value)}
                      required
                      minLength="200"
                    />
              </CFormGroup>
              <CFormGroup>
                <CLabel htmlFor="category">Category</CLabel>
                <CInput id="category" placeholder="Input Category" minLength="3" maxLength="100" value={category} onChange={(e)=>setCategory(e.target.value)} required/>
              </CFormGroup>
              <CButton color="primary" type="submit" onClick={()=>setStatus("Draft")}>
                Draft
              </CButton>
              <CButton color="info" className="ml-2" type="submit" onClick={()=>setStatus("Publish")}>
                Publish
              </CButton>
            </CForm>
            </CCardBody>
          </CCard>
        </CCol>
        </CRow>
    )
}

export default EditPost
