// src/App.js
import axios from 'axios'
import { useEffect, useState } from 'react'
import TestimonialForm from './TestimonialForm'

const App = () => {
  const [testimonials, setTestimonials] = useState([])
  const [selectedTestimonialId, setSelectedTestimonialId] = useState(null)

  useEffect(() => {
    fetchTestimonials()
  }, [])

  const fetchTestimonials = () => {
    axios
      .get('http://localhost:8080/testimonials')
      .then((response) => {
        setTestimonials(response.data)
      })
      .catch((error) => {
        console.error('There was an error fetching the testimonials!', error)
      })
  }

  const handleSuccess = () => {
    fetchTestimonials()
    setSelectedTestimonialId(null)
  }

  const handleDelete = (id) => {
    axios
      .delete(`http://localhost:8080/testimonials/${id}`)
      .then((response) => {
        console.log('Testimonial deleted successfully!', response)
        fetchTestimonials() // Refresh testimonials after deletion
      })
      .catch((error) => {
        console.error('There was an error deleting the testimonial!', error)
      })
  }

  return (
    <div className="container mx-auto px-4">
      <h1 className="text-3xl font-bold my-8">Testimonials</h1>
      <TestimonialForm
        testimonialId={selectedTestimonialId}
        onSuccess={handleSuccess}
      />
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mt-8">
        {testimonials.map((testimonial) => (
          <div
            key={testimonial.id}
            className="bg-white shadow-md rounded p-4"
          >
            <img
              src={`http://localhost:8080${testimonial.image}`}
              alt={testimonial.title}
              className="w-full h-48 object-cover rounded"
            />
            <h2 className="text-xl font-bold mt-4">{testimonial.title}</h2>
            <p className="text-gray-600 mt-2">{testimonial.description}</p>
            <p className="text-gray-500 mt-2">
              {new Date(testimonial.created_at).toLocaleDateString()}
            </p>
            <div className="mt-4">
              <button
                onClick={() => setSelectedTestimonialId(testimonial.id)}
                className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-700 mr-2"
              >
                Edit
              </button>
              <button
                onClick={() => handleDelete(testimonial.id)}
                className="bg-red-500 text-white py-2 px-4 rounded hover:bg-red-700"
              >
                Delete
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default App
