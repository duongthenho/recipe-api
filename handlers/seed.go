package handlers

import (
	"net/http"

	"recipe-api/db"
)

type seedStep struct {
	Text  string
	Image string
}

type seedRecipe struct {
	Title       string
	Image       string
	Cuisine     string
	Views       int
	Ingredients []string
	Steps       []seedStep
}

func SeedData(w http.ResponseWriter, r *http.Request) {

	data := []seedRecipe{
		{
			Title:   "Phở bò",
			Image:   "https://hongphatfood.com/wp-content/uploads/2020/05/vifon-beef-pho-6.jpg",
			Cuisine: "Việt Nam",
			Views:   1234,
			Ingredients: []string{
				"bánh phở", "thịt bò", "hành", "gừng", "quế",
			},
			Steps: []seedStep{
				{"Ninh xương bò lấy nước dùng", "https://file.hstatic.net/200000700229/article/cach-ham-xuong-bo-nhanh-mem_21eda1c3431544188a9cf0defc6113a8.jpg"},
				{"Nướng gừng và hành", "https://diwa.com.vn/wp-content/uploads/2024/08/cach-nau-pho-bo-bang-noi-ap-suat-2.jpg"},
				{"Cho gia vị vào nước dùng", "https://cafefcdn.com/203337114487263232/2025/2/3/image003-1738555036825838409020-1738565145106-17385651452681279816328.jpg"},
				{"Trụng bánh phở và xếp thịt bò", "https://nvnorthwest.com/wp-content/uploads/2021/11/qd1wgvr9-680x470.jpg"},
				{"Chan nước dùng và thưởng thức", "https://cdn2.tuoitre.vn/thumb_w/480/2019/10/26/qdi2985-15720670573741140523976.jpg"},
			},
		},
		{
			Title:   "Bún chả",
			Image:   "https://khaihoanphuquoc.com.vn/wp-content/uploads/2023/08/cach-lam-nuoc-mam-bun-cha-02.jpg",
			Cuisine: "Việt Nam",
			Views:   5231,
			Ingredients: []string{
				"bún", "thịt lợn", "nước mắm", "đu đủ", "tỏi",
			},
			Steps: []seedStep{
				{"Ướp thịt", "https://haiphu.vn/web/image/3525-268f086b/1-gia-vi-uop-thit-nuong.jpg?access_token=22098ab2-0a2f-471e-a51a-5e726417c60d"},
				{"Nướng thịt", "https://kungfu.com.vn/public/images/cach-uop-thit-nuong.jpg"},
				{"Pha nước chấm", "https://haiphu.vn/web/image/3350-40ff51e2/1-nuoc-cham-bun-thit-nuong.jpg?access_token=0160742c-d4a4-4a7b-96e9-9dd33db76dc9"},
				{"Ăn kèm bún và rau", "https://cdn.xanhsm.com/2025/01/40f46d34-bun-cha-ha-noi-o-tphcm-2-min.jpg"},
			},
		},
		{
			Title:   "Spaghetti Carbonara",
			Image:   "https://images.unsplash.com/photo-1603133872878-684f208fb84b",
			Cuisine: "Ý",
			Views:   151,
			Ingredients: []string{
				"mì spaghetti", "trứng", "phô mai", "thịt xông khói",
			},
			Steps: []seedStep{
				{"Luộc mì", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Chiên thịt xông khói", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Trộn trứng và phô mai", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Trộn với mì", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
			},
		},
		{
			Title:   "Sushi cuộn",
			Image:   "https://ibuki.vn/wp-content/uploads/2021/10/Com-cuon-dac-biet-sushi-scaled.jpeg",
			Cuisine: "Nhật Bản",
			Views:   255,
			Ingredients: []string{
				"cơm sushi", "rong biển", "cá hồi", "dưa leo",
			},
			Steps: []seedStep{
				{"Trải rong biển", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Cho cơm lên", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Xếp nhân", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Cuộn và cắt", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
			},
		},
		{
			Title:   "Cơm chiên trứng",
			Image:   "https://cdn11.dienmaycholon.vn/filewebdmclnew/public/userupload/files/kien-thuc/cach-lam-com-chien-trung/cach-lam-com-chien-trung-1.jpg",
			Cuisine: "Châu Á",
			Views:   23,
			Ingredients: []string{
				"cơm", "trứng", "hành lá", "nước mắm",
			},
			Steps: []seedStep{
				{"Phi hành", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Cho trứng", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Cho cơm", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
				{"Nêm gia vị", "https://vietgiao.edu.vn/wp-content/uploads/2021/04/hoc-nau-an-co-ban-1.jpg"},
			},
		},
	}

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for _, rcp := range data {

		res, err := tx.Exec(`
			INSERT INTO recipes(title, cuisine, image_url, views)
			VALUES (?, ?, ?, ?)
		`, rcp.Title, rcp.Cuisine, rcp.Image, rcp.Views)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), 500)
			return
		}

		recipeID, _ := res.LastInsertId()

		for _, ing := range rcp.Ingredients {
			_, err = tx.Exec(
				`INSERT INTO ingredients(recipe_id, name) VALUES (?, ?)`,
				recipeID, ing,
			)
			if err != nil {
				tx.Rollback()
				http.Error(w, err.Error(), 500)
				return
			}
		}

		for _, st := range rcp.Steps {
			_, err = tx.Exec(
				`INSERT INTO steps(recipe_id, text, image) VALUES (?, ?, ?)`,
				recipeID, st.Text, st.Image,
			)
			if err != nil {
				tx.Rollback()
				http.Error(w, err.Error(), 500)
				return
			}
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("seed ok"))
}