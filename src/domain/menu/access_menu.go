package menu

type AccessMenu struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	SubMenu []AccessSubMenu `json:"submenu"`
}

type AccessMenus []AccessMenu

type AccessSubMenu struct {
	ID       int    `json:"id"`
	MenuID   int    `json:"menu_id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Icon     string `json:"icon"`
	IsActive bool   `json:"is_active"`
}
