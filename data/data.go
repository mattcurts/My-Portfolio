package data

type TechLogo struct {
	Src   string
	Alt   string
	Title string
}

type WorkItem struct {
	Title     string
	Company   string
	Location  string
	Dates     string
	Logo      string
	Bullets   []string
	TechStack []TechLogo
}

type VolunteerItem struct {
	Title   string
	Company string
	Dates   string
	Logo    string
	Bullets []string
}

type ProjectItem struct {
	Title   string
	Bullets []string
}

// PageData contains all page data.
type PageData struct {
	WorkExperience      []WorkItem
	VolunteerExperience []VolunteerItem
	Projects            []ProjectItem
}

var Data = PageData{
	WorkExperience: []WorkItem{
		{
			Title:    "Full Stack Web Developer",
			Company:  "Good Things Consignment",
			Location: "Victoria, BC",
			Dates:    "Jan 2025 – Present",
			Logo:     "assets/company/virtool.svg",
			Bullets: []string{
				"Consignment & POS software to streamline store operations, replacing manual tracking with an automated inventory, sales, and payout system.",
				"Built full-stack iPad & web apps using AngularJS, Ionic, and an Express back-end with MongoDB, optimizing data flow and operational efficiency.",
				"Communicated technical challenges & pain points with users to refine features and usability—significantly reducing user errors.",
			},
			TechStack: []TechLogo{
				{Src: "assets/tech/AngularJS.svg", Alt: "AngularJS", Title: "AngularJS"},
				{Src: "assets/tech/HTML5.svg", Alt: "HTML5", Title: "HTML5"},
				{Src: "assets/tech/CSS3.svg", Alt: "CSS3", Title: "CSS3"},
				{Src: "assets/tech/Bootstrap.svg", Alt: "Bootstrap", Title: "Bootstrap"},
				{Src: "assets/tech/Ionic.svg", Alt: "Ionic", Title: "Ionic"},
				{Src: "assets/tech/MongoDB.svg", Alt: "MongoDB", Title: "MongoDB"},
				{Src: "assets/tech/Express.svg", Alt: "Express", Title: "Express"},
				{Src: "assets/tech/Node.js.svg", Alt: "Node.js", Title: "Node.js"},
			},
		},
		{
			Title:    "Backend Software Developer",
			Company:  "Virtool – CFIA",
			Location: "Remote",
			Dates:    "Sep 2023 – May 2024",
			Logo:     "assets/company/virtool.svg",
			Bullets: []string{
				"Contributed to Virtool, a bioinformatics platform for genomic analysis and pathogen tracking, supporting food safety and agricultural health.",
				"Implemented robust PyTest suites to ensure data integrity and system reliability across the platform.",
				"Led migration from MongoDB to SQL, enhancing data handling and query performance for large genomic datasets.",
			},
			TechStack: []TechLogo{
				{Src: "assets/tech/Python.svg", Alt: "Python", Title: "Python"},
				{Src: "assets/tech/Django.svg", Alt: "Django", Title: "Django"},
				{Src: "assets/tech/MongoDB.svg", Alt: "MongoDB", Title: "MongoDB"},
				{Src: "assets/tech/PostgresSQL.svg", Alt: "PostgreSQL", Title: "PostgreSQL"},
			},
		},
	},
	VolunteerExperience: []VolunteerItem{
		{
			Title:   "VP Corporate",
			Company: "UVic Engineering and Computer Science Student Society",
			Dates:   "May 2024 – May 2025",
			Logo:    "assets/company/ECSS.png",
			Bullets: []string{
				"Represented the council to Engineers and Geoscientists BC, ensuring alignment between industry standards and student interests.",
				"Established and maintained collaborative relationships with local industry, acting as the primary liaison between the society and industry partners.",
				"Secured sponsorships and industry participation for networking events while creating professional development opportunities for students.",
			},
		},
		{
			Title:   "President",
			Company: "UVic Engineering and Computer Science Student Society",
			Dates:   "Sep 2023 – May 2024",
			Logo:    "assets/company/ECSS.png",
			Bullets: []string{
				"Organized and facilitated semesterly general meetings with 60+ attendees, ensuring transparent governance and member participation.",
				"Streamlined study resources through digitizing the exam bank, boosting accessibility to various exams for improved studying experience.",
				"Raised over 4,500 dollars for the Vancouver Children's Health Fund by recruiting and organizing the volunteers for our charity event, Order of Pi.",
			},
		},
	},
	Projects: []ProjectItem{
		{
			Title: "Custom Video Codec",
			Bullets: []string{
				"Engineered a C++ compressor and decompressor for real-time decompression of 30 fps 480p video streams, ensuring seamless playback and improved user experience.",
				"Reduced YCbCr input file size by a factor of 4.5+, 12+, 14.8+ for low, medium, and high respectively – Achieving 1 bit per pixel for low and medium quality.",
				"Implemented P-frames and motion vector compensation techniques to optimize video compression, resulting in enhanced video quality and reduced file sizes.",
			},
		},
		{
			Title: "GZIP Compression Algorithm Implementation",
			Bullets: []string{
				"Created a C++ implementation of the GZIP compression algorithm, optimizing for speed and efficiency with support for DEFLATE block types 0, 1, and 2.",
				"Achieved an average of 200% compression of a 211 MB dataset in approximately 90 seconds on large files while maintaining memory efficiency.",
				"Utilized advanced data structures and algorithms, including LZSS, dynamic Huffman coding via the Package Merge Algorithm, and block type optimization logic.",
			},
		},
	},
}
