

(ns  assignment.core
  (:require [ring.adapter.jetty :as jetty]
            [ring.middleware.json :refer [wrap-json-response]]
            [ring.util.response :as response])
  (:import [java.sql DriverManager]))


(def students
  [{:id 1 :name "Aman" :subject "Clojure" :marks 85}
   {:id 2 :name "Rahul" :subject "Java" :marks 45}
   {:id 3 :name "Priya" :subject "Clojure" :marks 92}
   {:id 4 :name "Neha" :subject "Python" :marks 35}
   {:id 5 :name "Rohit" :subject "Java" :marks 76}])


(defn add-grade [student]
  (assoc student :grade
         (cond
           (>= (:marks student) 90) "A"
           (>= (:marks student) 75) "B"
           (>= (:marks student) 50) "C"
           :else "F")))

;; FILTER: Select passed students
(defn passed-students [data]
  (filter #(>= (:marks %) 40) data))

;; REDUCE: Calculate statistics
(defn calculate-summary [data]
  (reduce
   (fn [result student]
     (-> result
         (update :total-marks + (:marks student))
         (update :total-students inc)))
   {:total-marks 0 :total-students 0}
   data))

(def processed-students
  (map add-grade students))

(def passed
  (passed-students processed-students))

(def summary
  (let [stats (calculate-summary processed-students)]
    (assoc stats :average
           (if (pos? (:total-students stats))
             (/ (:total-marks stats)
                (:total-students stats))
             0))))

;; --------------------------------
;; 3. Database Connection
;; --------------------------------

(def db-url "jdbc:sqlite:student.db")

(defn get-connection []
  (DriverManager/getConnection db-url))

;; --------------------------------
;; 4. Create Database Table
;; --------------------------------

(defn create-table []
  (try
    (with-open [conn (get-connection)
                stmt (.createStatement conn)]
      (.executeUpdate stmt
                      "CREATE TABLE IF NOT EXISTS students
                                           (id INTEGER PRIMARY KEY,
                                            name TEXT,
                                      subject TEXT,
                                  marks INTEGER,
                                  grade TEXT)")
      (println "Database table created successfully"))

    (catch Exception e
      (println "Database error:" (.getMessage e)))))

;; --------------------------------
;; 5. Insert Student Records
;; --------------------------------

(defn insert-students [data]
  (try
    (with-open [conn (get-connection)
                stmt (.prepareStatement conn
                                        "INSERT OR REPLACE INTO students
                                                                                                     (id, name, subject, marks, grade)
                                                                 VALUES (?, ?, ?, ?, ?)")]

      (doseq [student data]
        (.setInt stmt 1 (:id student))
        (.setString stmt 2 (:name student))
        (.setString stmt 3 (:subject student))
        (.setInt stmt 4 (:marks student))
        (.setString stmt 5 (:grade student))
        (.addBatch stmt))

      (.executeBatch stmt)
      (println "Student records inserted successfully"))

    (catch Exception e
      (println "Insert error:" (.getMessage e)))))

;; --------------------------------
;; 6. Read Student Records
;; --------------------------------

(defn get-students []
  (try
    (with-open [conn (get-connection)
                stmt (.createStatement conn)
                rs (.executeQuery stmt
                                  "SELECT * FROM students")]

      (loop [result []]
        (if (.next rs)
          (recur
           (conj result
                 {:id (.getInt rs "id")
                  :name (.getString rs "name")
                  :subject (.getString rs "subject")
                  :marks (.getInt rs "marks")
                  :grade (.getString rs "grade")}))
          result)))

    (catch Exception e
      (println "Read error:" (.getMessage e))
      [])))

;; --------------------------------
;; 7. Ring HTTP Endpoints
;; --------------------------------

(defn app [request]
  (case (:uri request)

    "/students"
    (response/response (get-students))

    "/passed"
    (response/response (vec (passed-students (get-students))))

    "/summary"
    (response/response
     (let [data (get-students)
           stats (calculate-summary data)]
       (assoc stats :average
              (if (pos? (:total-students stats))
                (/ (:total-marks stats)
                   (:total-students stats))
                0))))

    "/health"
    (response/response {:status "Application running"})

    (response/not-found {:error "Endpoint not found"})))

(def app-json
  (wrap-json-response app))

;; --------------------------------
;; 8. Main Function
;; --------------------------------

(defn -main []
  (create-table)
  (insert-students processed-students)

  (println "\nProcessed Student Data:")
  (doseq [student processed-students]
    (println student))

  (println "\nSummary:")
  (println summary)

  (println "\nStarting Ring Server on port 3000...")

  (jetty/run-jetty app-json
                   {:port 3000
                    :join? true}))
