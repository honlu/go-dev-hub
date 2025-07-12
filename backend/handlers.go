package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 分页结构
type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

// 管理接口 - 博客 CRUD
func createBlog(c *gin.Context) {
	var blog Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误", Data: nil})
		return
	}

	if err := db.Create(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "创建失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "创建成功", Data: blog})
}

func getBlogs(c *gin.Context) {
	var blogs []Blog
	var total int64

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	// 查询条件
	query := db.Model(&Blog{})
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if title := c.Query("title"); title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	// 获取总数
	query.Count(&total)

	// 获取数据
	if err := query.Offset(offset).Limit(pageSize).Find(&blogs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "查询成功",
		Data: gin.H{
			"blogs": blogs,
			"pagination": Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    total,
			},
		},
	})
}

func getBlog(c *gin.Context) {
	id := c.Param("id")
	var blog Blog

	if err := db.First(&blog, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, Response{Code: 404, Message: "博客不存在", Data: nil})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		}
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "查询成功", Data: blog})
}

func updateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog Blog

	if err := db.First(&blog, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, Response{Code: 404, Message: "博客不存在", Data: nil})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		}
		return
	}

	var updateData Blog
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误", Data: nil})
		return
	}

	if err := db.Model(&blog).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "更新失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "更新成功", Data: blog})
}

func deleteBlog(c *gin.Context) {
	id := c.Param("id")
	var blog Blog

	if err := db.First(&blog, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, Response{Code: 404, Message: "博客不存在", Data: nil})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		}
		return
	}

	if err := db.Delete(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "删除失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "删除成功", Data: nil})
}

// 管理接口 - 题库 CRUD
func createQuestion(c *gin.Context) {
	var question Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误", Data: nil})
		return
	}

	if err := db.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "创建失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "创建成功", Data: question})
}

func getQuestions(c *gin.Context) {
	var questions []Question
	var total int64

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	// 查询条件
	query := db.Model(&Question{})
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}
	if difficulty := c.Query("difficulty"); difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}
	if title := c.Query("title"); title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	// 获取总数
	query.Count(&total)

	// 获取数据
	if err := query.Offset(offset).Limit(pageSize).Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "查询成功",
		Data: gin.H{
			"questions": questions,
			"pagination": Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    total,
			},
		},
	})
}

func getQuestion(c *gin.Context) {
	id := c.Param("id")
	var question Question

	if err := db.First(&question, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, Response{Code: 404, Message: "题目不存在", Data: nil})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		}
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "查询成功", Data: question})
}

func updateQuestion(c *gin.Context) {
	id := c.Param("id")
	var question Question

	if err := db.First(&question, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, Response{Code: 404, Message: "题目不存在", Data: nil})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		}
		return
	}

	var updateData Question
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误", Data: nil})
		return
	}

	if err := db.Model(&question).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "更新失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "更新成功", Data: question})
}

func deleteQuestion(c *gin.Context) {
	id := c.Param("id")
	var question Question

	if err := db.First(&question, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, Response{Code: 404, Message: "题目不存在", Data: nil})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		}
		return
	}

	if err := db.Delete(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "删除失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "删除成功", Data: nil})
}

// 只读接口 - 博客查询
func getPublicBlogs(c *gin.Context) {
	var blogs []Blog
	var total int64

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	// 只查询已发布的博客
	query := db.Model(&Blog{}).Where("status = ?", "published")
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}
	if title := c.Query("title"); title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	// 获取总数
	query.Count(&total)

	// 获取数据
	if err := query.Offset(offset).Limit(pageSize).Find(&blogs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "查询成功",
		Data: gin.H{
			"blogs": blogs,
			"pagination": Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    total,
			},
		},
	})
}

func getPublicBlog(c *gin.Context) {
	id := c.Param("id")
	var blog Blog

	if err := db.Where("id = ? AND status = ?", id, "published").First(&blog).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, Response{Code: 404, Message: "博客不存在", Data: nil})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		}
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "查询成功", Data: blog})
}

// 只读接口 - 题库查询
func getPublicQuestions(c *gin.Context) {
	var questions []Question
	var total int64

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	// 查询条件
	query := db.Model(&Question{})
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}
	if difficulty := c.Query("difficulty"); difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}
	if title := c.Query("title"); title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	// 获取总数
	query.Count(&total)

	// 获取数据
	if err := query.Offset(offset).Limit(pageSize).Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "查询成功",
		Data: gin.H{
			"questions": questions,
			"pagination": Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    total,
			},
		},
	})
}

func getPublicQuestion(c *gin.Context) {
	id := c.Param("id")
	var question Question

	if err := db.First(&question, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, Response{Code: 404, Message: "题目不存在", Data: nil})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败", Data: nil})
		}
		return
	}

	c.JSON(http.StatusOK, Response{Code: 200, Message: "查询成功", Data: question})
}

// 获取分类枚举
func getCategories(c *gin.Context) {
	// 统一分类 - Go全栈架构师相关
	categories := []string{
		"Go开发", "系统架构", "微服务", "云原生", "数据库", "前端技术", "DevOps", "性能优化", "算法", "其他",
	}

	// 题目难度
	difficulties := []map[string]string{
		{"label": "简单", "value": "easy"},
		{"label": "中等", "value": "medium"},
		{"label": "困难", "value": "hard"},
	}

	// 分类对应的标签映射 - 统一标签体系
	categoryTags := map[string][]string{
		// Go开发相关标签
		"Go开发": {"Golang", "Gin", "Echo", "GORM", "Goroutine", "Channel", "并发编程", "性能优化", "单元测试", "Web框架"},

		// 系统架构相关标签
		"系统架构": {"分布式", "高并发", "高可用", "可扩展性", "负载均衡", "缓存", "消息队列", "监控", "故障排查", "架构设计"},

		// 微服务相关标签
		"微服务": {"服务拆分", "服务治理", "服务发现", "负载均衡", "熔断降级", "链路追踪", "API网关", "分布式事务", "服务网格", "Istio"},

		// 云原生相关标签
		"云原生": {"Docker", "Kubernetes", "Istio", "Prometheus", "Grafana", "服务网格", "容器化", "可观测性", "弹性伸缩", "Helm"},

		// 数据库相关标签
		"数据库": {"MySQL", "PostgreSQL", "Redis", "MongoDB", "SQL", "索引优化", "事务", "分库分表", "性能调优", "监控告警"},

		// 前端技术相关标签
		"前端技术": {"JavaScript", "TypeScript", "Vue", "React", "HTML", "CSS", "Webpack", "组件化", "状态管理", "UI框架"},

		// DevOps相关标签
		"DevOps": {"Jenkins", "GitLab CI", "Docker", "Kubernetes", "Terraform", "监控告警", "自动化部署", "持续集成", "持续部署", "日志收集"},

		// 性能优化相关标签
		"性能优化": {"性能调优", "内存优化", "CPU优化", "网络优化", "数据库优化", "缓存优化", "代码优化", "架构优化", "监控分析", "瓶颈分析"},

		// 算法相关标签
		"算法": {"排序算法", "查找算法", "动态规划", "贪心算法", "回溯算法", "分治算法", "图论算法", "字符串算法", "数学算法", "复杂度分析"},

		// 其他分类的通用标签
		"其他": {"技术分享", "经验总结", "问题解决", "最佳实践", "技术趋势", "开源项目", "工具推荐", "学习资源", "技术社区", "行业动态"},
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取成功",
		Data: gin.H{
			"categories":    categories,
			"difficulties":  difficulties,
			"category_tags": categoryTags,
		},
	})
}
