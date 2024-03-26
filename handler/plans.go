package handler

func PlansHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllPlans(w,r)

		w.Header().Set("Content-Type", "application/json")
		fmt.Println("'GET'-response sent to /plans on", time.Now().Format(time.RFC850))
	case "POST":
		createPlan(w,r)

		reqMessage := "'POST'-request sent."
		w.Header().Set("Content-Type", "application/json") 
		fmt.Fprintf(w,`{"request": "%s"}`, reqMessage)
		fmt.Println("'POST'-response sent to /plans on", time.Now().Format(time.RFC850))
	case "PUT":
		updatePlan(w,r)

		reqMessage := "'PUT'-request sent."
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w,`{"request": "%s"}`, reqMessage)
		fmt.Println("'PUT'-response sent to /plans on", time.Now().Format(time.RFC850))
	case "DELETE":
		deletePlan(w,r)	
		
		reqMessage := "'DELETE'-request sent."
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w,`{"request": "%s"}`, reqMessage)
		fmt.Println("'DELETE'-response sent to /plans on", time.Now().Format(time.RFC850))
	}
}

func getAllPlans(w http.ResponseWriter, r *http.Request) {
	/*Something*/
}
func createPlan(w http.ResponseWriter, r *http.Request) {
	/*Something*/
}
func updatePlan(w http.ResponseWriter, r *http.Request) {
	/*Something*/
}
func deletePlan(w http.ResponseWriter, r *http.Request) {
	/*Something*/
}
