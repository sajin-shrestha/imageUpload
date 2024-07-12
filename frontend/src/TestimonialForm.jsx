// src/components/TestimonialForm.js
import axios from 'axios'
import { useEffect, useState } from 'react'

const TestimonialForm = ({ testimonialId, onSuccess }) => {
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [image, setImage] = useState(null)

  useEffect(() => {
    if (testimonialId) {
      axios
        .get(`http://localhost:8080/testimonials/${testimonialId}`)
        .then((response) => {
          const { title, description, image } = response.data
          setTitle(title)
          setDescription(description)
          setImage(image)
        })
        .catch((error) => {
          console.error('There was an error fetching the testimonial!', error)
        })
    }
  }, [testimonialId])

  const handleSubmit = (e) => {
    e.preventDefault()

    const formData = new FormData()
    formData.append('title', title)
    formData.append('description', description)
    if (image) formData.append('image', image)

    const request = testimonialId
      ? axios.put(
          `http://localhost:8080/testimonials/${testimonialId}`,
          formData,
        )
      : axios.post('http://localhost:8080/testimonials', formData)

    request
      .then((response) => {
        onSuccess(response.data)
      })
      .catch((error) => {
        console.error('There was an error submitting the testimonial!', error)
      })
  }

  return (
    <form
      onSubmit={handleSubmit}
      className="max-w-md mx-auto mt-8"
    >
      <div className="mb-4">
        <label
          htmlFor="title"
          className="block text-gray-700"
        >
          Title
        </label>
        <input
          type="text"
          id="title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="w-full px-3 py-2 border rounded"
          required
        />
      </div>
      <div className="mb-4">
        <label
          htmlFor="description"
          className="block text-gray-700"
        >
          Description
        </label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="w-full px-3 py-2 border rounded"
          required
        />
      </div>
      <div className="mb-4">
        <label
          htmlFor="image"
          className="block text-gray-700"
        >
          Image
        </label>
        <input
          type="file"
          id="image"
          onChange={(e) => setImage(e.target.files[0])}
          className="w-full px-3 py-2 border rounded"
          accept="image/*"
        />
      </div>
      {image && (
        <div className="mb-4">
          {/* Preview the selected image */}
          <img
            src={URL.createObjectURL(image)}
            alt="Preview"
            className="w-full h-40 object-cover rounded mb-2"
          />
        </div>
      )}
      <button
        type="submit"
        className="w-full bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-700"
      >
        {testimonialId ? 'Update Testimonial' : 'Create Testimonial'}
      </button>
    </form>
  )
}

export default TestimonialForm
